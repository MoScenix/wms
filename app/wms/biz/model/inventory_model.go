package model

import (
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"gorm.io/gorm"
)

type Inventory struct {
	gorm.Model
	WarehouseID       uint   `gorm:"not null;index:idx_inv_bucket,unique"`
	LocationID        uint   `gorm:"not null;index:idx_inv_bucket,unique"`
	SkuID             uint   `gorm:"not null;index:idx_inv_bucket,unique"`
	LotNo             string `gorm:"type:varchar(64);not null;default:'';index:idx_inv_bucket,unique"`
	ExpireDate        string `gorm:"type:varchar(32);not null;default:'';index:idx_inv_bucket,unique"`
	Quantity          int64  `gorm:"not null;default:0"`
	LockedQuantity    int64  `gorm:"not null;default:0"`
	AvailableQuantity int64  `gorm:"not null;default:0"`
}

func (Inventory) TableName() string { return "inventory" }

func toInventoryInfo(m Inventory) *wms.InventoryInfo {
	return &wms.InventoryInfo{Id: int64(m.ID), WarehouseId: int64(m.WarehouseID), LocationId: int64(m.LocationID), SkuId: int64(m.SkuID), LotNo: m.LotNo, ExpireDate: m.ExpireDate, Quantity: m.Quantity, LockedQuantity: m.LockedQuantity, AvailableQuantity: m.AvailableQuantity, CreateTime: toTimeString(m.CreatedAt), UpdateTime: toTimeString(m.UpdatedAt)}
}

func (p *WmsProQuery) GetInventory(id int64) (*wms.GetInventoryResp, error) {
	cacheKey := p.key("inventory", id)
	if cached := new(wms.InventoryInfo); p.cacheGet(cacheKey, cached) {
		return &wms.GetInventoryResp{Inventory: cached}, nil
	}
	var m Inventory
	if err := p.q.db.WithContext(p.q.ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	data := toInventoryInfo(m)
	p.cacheSet(cacheKey, data)
	return &wms.GetInventoryResp{Inventory: data}, nil
}

func (p *WmsProQuery) ListInventory(req *wms.ListInventoryReq) (*wms.ListInventoryResp, error) {
	pageNum, pageSize := pageNormalize(req.PageNum, req.PageSize)
	q := p.q.db.WithContext(p.q.ctx).Model(&Inventory{})
	if req.WarehouseId > 0 {
		q = q.Where("warehouse_id = ?", req.WarehouseId)
	}
	if req.LocationId > 0 {
		q = q.Where("location_id = ?", req.LocationId)
	}
	if req.SkuId > 0 {
		q = q.Where("sku_id = ?", req.SkuId)
	}
	if req.LotNo != "" {
		q = q.Where("lot_no like ?", likePrefix(req.LotNo))
	}
	if req.ExpireDate != "" {
		q = q.Where("expire_date like ?", likePrefix(req.ExpireDate))
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}
	var rows []Inventory
	if err := q.Order("id desc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}
	ret := make([]*wms.InventoryInfo, 0, len(rows))
	for _, it := range rows {
		ret = append(ret, toInventoryInfo(it))
	}
	return &wms.ListInventoryResp{Records: ret, Total: total}, nil
}

func (p *WmsProQuery) CheckInventory(req *wms.CheckInventoryReq) (*wms.CheckInventoryResp, error) {
	resp := &wms.CheckInventoryResp{Results: make([]*wms.CheckInventoryResultItem, 0, len(req.Items)), AllPass: true}
	for _, item := range req.Items {
		result := wms.InventoryCheckResult_CHECK_OK
		available := int64(0)

		var skuCount int64
		if err := p.q.db.WithContext(p.q.ctx).Model(&Sku{}).Where("id = ?", item.SkuId).Count(&skuCount).Error; err != nil {
			return nil, err
		}
		if skuCount == 0 {
			result = wms.InventoryCheckResult_CHECK_SKU_NOT_FOUND
		} else {
			var locCount int64
			if err := p.q.db.WithContext(p.q.ctx).Model(&Location{}).Where("id = ?", item.LocationId).Count(&locCount).Error; err != nil {
				return nil, err
			}
			if locCount == 0 {
				result = wms.InventoryCheckResult_CHECK_LOCATION_NOT_FOUND
			} else {
				var inv Inventory
				err := p.q.db.WithContext(p.q.ctx).Where("warehouse_id = ? and location_id = ? and sku_id = ? and lot_no = ? and expire_date = ?",
					item.WarehouseId, item.LocationId, item.SkuId, item.LotNo, item.ExpireDate).First(&inv).Error
				if err == nil {
					available = inv.AvailableQuantity
				}
				if available < item.RequiredQty {
					result = wms.InventoryCheckResult_CHECK_INSUFFICIENT
				}
			}
		}
		if result != wms.InventoryCheckResult_CHECK_OK {
			resp.AllPass = false
		}
		resp.Results = append(resp.Results, &wms.CheckInventoryResultItem{Requested: item, Result: result, AvailableQty: available})
	}
	return resp, nil
}

package model

import (
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"gorm.io/gorm"
)

type Warehouse struct {
	gorm.Model
	WarehouseCode string `gorm:"type:varchar(64);not null;uniqueIndex"`
	WarehouseName string `gorm:"type:varchar(128);not null"`
	Address       string `gorm:"type:varchar(255);default:''"`
	ManagerUserID int64  `gorm:"not null;default:0"`
}

func (Warehouse) TableName() string { return "warehouse" }

func toWarehouseInfo(m Warehouse) *wms.WarehouseInfo {
	return &wms.WarehouseInfo{
		Id:            int64(m.ID),
		WarehouseCode: m.WarehouseCode,
		WarehouseName: m.WarehouseName,
		Address:       m.Address,
		ManagerUserId: m.ManagerUserID,
		CreateTime:    toTimeString(m.CreatedAt),
		UpdateTime:    toTimeString(m.UpdatedAt),
	}
}

func (p *WmsProQuery) AddWarehouse(req *wms.AddWarehouseReq) (*wms.AddWarehouseResp, error) {
	m := Warehouse{WarehouseCode: req.WarehouseCode, WarehouseName: req.WarehouseName, Address: req.Address, ManagerUserID: req.ManagerUserId}
	if err := p.q.db.WithContext(p.q.ctx).Create(&m).Error; err != nil {
		return nil, err
	}
	return &wms.AddWarehouseResp{Id: int64(m.ID)}, nil
}

func (p *WmsProQuery) UpdateWarehouse(req *wms.UpdateWarehouseReq) (*wms.UpdateWarehouseResp, error) {
	updates := map[string]any{}
	if req.WarehouseName != "" {
		updates["warehouse_name"] = req.WarehouseName
	}
	if req.Address != "" {
		updates["address"] = req.Address
	}
	if req.ManagerUserId > 0 {
		updates["manager_user_id"] = req.ManagerUserId
	}
	if len(updates) == 0 {
		return &wms.UpdateWarehouseResp{Success: true}, nil
	}
	if err := p.q.db.WithContext(p.q.ctx).Model(&Warehouse{}).Where("id = ?", req.Id).Updates(updates).Error; err != nil {
		return nil, err
	}
	p.cacheDel(p.key("warehouse", req.Id))
	return &wms.UpdateWarehouseResp{Success: true}, nil
}

func (p *WmsProQuery) GetWarehouse(id int64) (*wms.GetWarehouseResp, error) {
	cacheKey := p.key("warehouse", id)
	if cached := new(wms.WarehouseInfo); p.cacheGet(cacheKey, cached) {
		return &wms.GetWarehouseResp{Warehouse: cached}, nil
	}
	var m Warehouse
	if err := p.q.db.WithContext(p.q.ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	data := toWarehouseInfo(m)
	p.cacheSet(cacheKey, data)
	return &wms.GetWarehouseResp{Warehouse: data}, nil
}

func (p *WmsProQuery) ListWarehouse(req *wms.ListWarehouseReq) (*wms.ListWarehouseResp, error) {
	pageNum, pageSize := pageNormalize(req.PageNum, req.PageSize)
	q := p.q.db.WithContext(p.q.ctx).Model(&Warehouse{}).
		Where("warehouse_code like ?", likePrefix(req.WarehouseCode)).
		Where("warehouse_name like ?", likePrefix(req.WarehouseName))
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}
	var rows []Warehouse
	if err := q.Order("id desc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}
	ret := make([]*wms.WarehouseInfo, 0, len(rows))
	for _, it := range rows {
		ret = append(ret, toWarehouseInfo(it))
	}
	return &wms.ListWarehouseResp{Records: ret, Total: total}, nil
}

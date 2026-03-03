package model

import (
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"gorm.io/gorm"
)

type InventoryTxn struct {
	gorm.Model
	BizType        int32  `gorm:"not null;index"`
	BizNo          string `gorm:"type:varchar(64);not null;index"`
	WarehouseID    uint   `gorm:"not null;index"`
	LocationID     uint   `gorm:"not null;index"`
	SkuID          uint   `gorm:"not null;index"`
	LotNo          string `gorm:"type:varchar(64);default:''"`
	ExpireDate     string `gorm:"type:varchar(32);default:''"`
	QtyChange      int64  `gorm:"not null"`
	BeforeQty      int64  `gorm:"not null"`
	AfterQty       int64  `gorm:"not null"`
	OperatorUserID int64  `gorm:"not null;default:0"`
}

func (InventoryTxn) TableName() string { return "inventory_txn" }

func (p *WmsProQuery) ListInventoryTxn(req *wms.ListInventoryTxnReq) (*wms.ListInventoryTxnResp, error) {
	pageNum, pageSize := pageNormalize(req.PageNum, req.PageSize)
	q := p.q.db.WithContext(p.q.ctx).Model(&InventoryTxn{})
	if req.BizType == wms.BizType_OUTBOUND || req.BizType == wms.BizType_ADJUST {
		q = q.Where("biz_type = ?", int32(req.BizType))
	}
	if req.BizNo != "" {
		q = q.Where("biz_no like ?", likePrefix(req.BizNo))
	}
	if req.WarehouseId > 0 {
		q = q.Where("warehouse_id = ?", req.WarehouseId)
	}
	if req.LocationId > 0 {
		q = q.Where("location_id = ?", req.LocationId)
	}
	if req.SkuId > 0 {
		q = q.Where("sku_id = ?", req.SkuId)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}
	var rows []InventoryTxn
	if err := q.Order("id desc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}
	ret := make([]*wms.InventoryTxnInfo, 0, len(rows))
	for _, row := range rows {
		ret = append(ret, &wms.InventoryTxnInfo{Id: int64(row.ID), BizType: wms.BizType(row.BizType), BizNo: row.BizNo, WarehouseId: int64(row.WarehouseID), LocationId: int64(row.LocationID), SkuId: int64(row.SkuID), LotNo: row.LotNo, ExpireDate: row.ExpireDate, QtyChange: row.QtyChange, BeforeQty: row.BeforeQty, AfterQty: row.AfterQty, OperatorUserId: row.OperatorUserID, CreateTime: toTimeString(row.CreatedAt)})
	}
	return &wms.ListInventoryTxnResp{Records: ret, Total: total}, nil
}

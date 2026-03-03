package model

import (
	"fmt"
	"time"

	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type InboundOrder struct {
	gorm.Model
	InboundNo      string `gorm:"type:varchar(64);not null;uniqueIndex"`
	WarehouseID    uint   `gorm:"not null;index"`
	SupplierName   string `gorm:"type:varchar(128);default:''"`
	Status         int32  `gorm:"not null;default:0"`
	OperatorUserID int64  `gorm:"not null;default:0"`
	Remark         string `gorm:"type:varchar(255);default:''"`
}

func (InboundOrder) TableName() string { return "inbound_order" }

type InboundOrderItem struct {
	gorm.Model
	InboundOrderID uint   `gorm:"not null;index"`
	SkuID          uint   `gorm:"not null;index"`
	LocationID     uint   `gorm:"not null;index"`
	LotNo          string `gorm:"type:varchar(64);default:''"`
	ExpireDate     string `gorm:"type:varchar(32);default:''"`
	Qty            int64  `gorm:"not null"`
}

func (InboundOrderItem) TableName() string { return "inbound_order_item" }

func toInboundItemInfo(m InboundOrderItem) *wms.InboundOrderItemInfo {
	return &wms.InboundOrderItemInfo{Id: int64(m.ID), InboundOrderId: int64(m.InboundOrderID), SkuId: int64(m.SkuID), LocationId: int64(m.LocationID), LotNo: m.LotNo, ExpireDate: m.ExpireDate, Qty: m.Qty}
}

func (p *WmsProQuery) CreateInboundOrder(req *wms.CreateInboundOrderReq) (*wms.CreateInboundOrderResp, error) {
	var orderID int64
	err := p.q.db.WithContext(p.q.ctx).Transaction(func(tx *gorm.DB) error {
		order := InboundOrder{InboundNo: fmt.Sprintf("INB%d", time.Now().UnixNano()), WarehouseID: uint(req.WarehouseId), SupplierName: req.SupplierName, Status: int32(wms.InboundStatus_INBOUND_DRAFT), OperatorUserID: req.OperatorUserId, Remark: req.Remark}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		for _, it := range req.Items {
			item := InboundOrderItem{InboundOrderID: order.ID, SkuID: uint(it.SkuId), LocationID: uint(it.LocationId), LotNo: it.LotNo, ExpireDate: it.ExpireDate, Qty: it.Qty}
			if err := tx.Create(&item).Error; err != nil {
				return err
			}
		}
		orderID = int64(order.ID)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &wms.CreateInboundOrderResp{Id: orderID}, nil
}

func (p *WmsProQuery) ReceiveInboundOrder(req *wms.ReceiveInboundOrderReq) (*wms.ReceiveInboundOrderResp, error) {
	err := p.q.db.WithContext(p.q.ctx).Transaction(func(tx *gorm.DB) error {
		var order InboundOrder
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&order, req.Id).Error; err != nil {
			return err
		}
		if order.Status == int32(wms.InboundStatus_INBOUND_RECEIVED) {
			return nil
		}
		if order.Status == int32(wms.InboundStatus_INBOUND_CANCELED) {
			return fmt.Errorf("inbound order canceled")
		}

		var items []InboundOrderItem
		if err := tx.Where("inbound_order_id = ?", order.ID).Find(&items).Error; err != nil {
			return err
		}
		for _, it := range items {
			var inv Inventory
			err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
				Where("warehouse_id = ? and location_id = ? and sku_id = ? and lot_no = ? and expire_date = ?", order.WarehouseID, it.LocationID, it.SkuID, it.LotNo, it.ExpireDate).
				First(&inv).Error
			if err != nil {
				if err == gorm.ErrRecordNotFound {
					inv = Inventory{WarehouseID: order.WarehouseID, LocationID: it.LocationID, SkuID: it.SkuID, LotNo: it.LotNo, ExpireDate: it.ExpireDate}
					if err := tx.Create(&inv).Error; err != nil {
						return err
					}
				} else {
					return err
				}
			}
			before := inv.Quantity
			inv.Quantity += it.Qty
			inv.AvailableQuantity = inv.Quantity - inv.LockedQuantity
			if err := tx.Model(&Inventory{}).Where("id = ?", inv.ID).Updates(map[string]any{"quantity": inv.Quantity, "available_quantity": inv.AvailableQuantity}).Error; err != nil {
				return err
			}

			txn := InventoryTxn{BizType: int32(wms.BizType_INBOUND), BizNo: order.InboundNo, WarehouseID: order.WarehouseID, LocationID: it.LocationID, SkuID: it.SkuID, LotNo: it.LotNo, ExpireDate: it.ExpireDate, QtyChange: it.Qty, BeforeQty: before, AfterQty: inv.Quantity, OperatorUserID: req.OperatorUserId}
			if err := tx.Create(&txn).Error; err != nil {
				return err
			}
			p.cacheDel(p.key("inventory", int64(inv.ID)))
		}
		if err := tx.Model(&InboundOrder{}).Where("id = ?", order.ID).Updates(map[string]any{"status": int32(wms.InboundStatus_INBOUND_RECEIVED), "operator_user_id": req.OperatorUserId}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &wms.ReceiveInboundOrderResp{Success: true}, nil
}

func (p *WmsProQuery) GetInboundOrder(id int64) (*wms.GetInboundOrderResp, error) {
	var order InboundOrder
	if err := p.q.db.WithContext(p.q.ctx).First(&order, id).Error; err != nil {
		return nil, err
	}
	var items []InboundOrderItem
	if err := p.q.db.WithContext(p.q.ctx).Where("inbound_order_id = ?", order.ID).Find(&items).Error; err != nil {
		return nil, err
	}
	itemList := make([]*wms.InboundOrderItemInfo, 0, len(items))
	for _, it := range items {
		itemList = append(itemList, toInboundItemInfo(it))
	}
	return &wms.GetInboundOrderResp{Order: &wms.InboundOrderInfo{Id: int64(order.ID), InboundNo: order.InboundNo, WarehouseId: int64(order.WarehouseID), SupplierName: order.SupplierName, Status: wms.InboundStatus(order.Status), OperatorUserId: order.OperatorUserID, Remark: order.Remark, CreateTime: toTimeString(order.CreatedAt), UpdateTime: toTimeString(order.UpdatedAt), Items: itemList}}, nil
}

func (p *WmsProQuery) ListInboundOrder(req *wms.ListInboundOrderReq) (*wms.ListInboundOrderResp, error) {
	pageNum, pageSize := pageNormalize(req.PageNum, req.PageSize)
	q := p.q.db.WithContext(p.q.ctx).Model(&InboundOrder{})
	if req.WarehouseId > 0 {
		q = q.Where("warehouse_id = ?", req.WarehouseId)
	}
	if req.InboundNo != "" {
		q = q.Where("inbound_no like ?", likePrefix(req.InboundNo))
	}
	if req.Status == wms.InboundStatus_INBOUND_RECEIVED || req.Status == wms.InboundStatus_INBOUND_CANCELED {
		q = q.Where("status = ?", int32(req.Status))
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}
	var rows []InboundOrder
	if err := q.Order("id desc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}
	ret := make([]*wms.InboundOrderInfo, 0, len(rows))
	for _, row := range rows {
		var items []InboundOrderItem
		if err := p.q.db.WithContext(p.q.ctx).Where("inbound_order_id = ?", row.ID).Find(&items).Error; err != nil {
			return nil, err
		}
		itemList := make([]*wms.InboundOrderItemInfo, 0, len(items))
		for _, it := range items {
			itemList = append(itemList, toInboundItemInfo(it))
		}
		ret = append(ret, &wms.InboundOrderInfo{Id: int64(row.ID), InboundNo: row.InboundNo, WarehouseId: int64(row.WarehouseID), SupplierName: row.SupplierName, Status: wms.InboundStatus(row.Status), OperatorUserId: row.OperatorUserID, Remark: row.Remark, CreateTime: toTimeString(row.CreatedAt), UpdateTime: toTimeString(row.UpdatedAt), Items: itemList})
	}
	return &wms.ListInboundOrderResp{Records: ret, Total: total}, nil
}

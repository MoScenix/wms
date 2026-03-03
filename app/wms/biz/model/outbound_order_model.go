package model

import (
	"fmt"
	"time"

	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type OutboundOrder struct {
	gorm.Model
	OutboundNo     string `gorm:"type:varchar(64);not null;uniqueIndex"`
	WarehouseID    uint   `gorm:"not null;index"`
	CustomerName   string `gorm:"type:varchar(128);default:''"`
	Status         int32  `gorm:"not null;default:0"`
	OperatorUserID int64  `gorm:"not null;default:0"`
	Remark         string `gorm:"type:varchar(255);default:''"`
}

func (OutboundOrder) TableName() string { return "outbound_order" }

type OutboundOrderItem struct {
	gorm.Model
	OutboundOrderID uint   `gorm:"not null;index"`
	SkuID           uint   `gorm:"not null;index"`
	LocationID      uint   `gorm:"not null;index"`
	LotNo           string `gorm:"type:varchar(64);default:''"`
	ExpireDate      string `gorm:"type:varchar(32);default:''"`
	Qty             int64  `gorm:"not null"`
}

func (OutboundOrderItem) TableName() string { return "outbound_order_item" }

func toOutboundItemInfo(m OutboundOrderItem) *wms.OutboundOrderItemInfo {
	return &wms.OutboundOrderItemInfo{Id: int64(m.ID), OutboundOrderId: int64(m.OutboundOrderID), SkuId: int64(m.SkuID), LocationId: int64(m.LocationID), LotNo: m.LotNo, ExpireDate: m.ExpireDate, Qty: m.Qty}
}

func (p *WmsProQuery) CreateOutboundOrder(req *wms.CreateOutboundOrderReq) (*wms.CreateOutboundOrderResp, error) {
	var orderID int64
	err := p.q.db.WithContext(p.q.ctx).Transaction(func(tx *gorm.DB) error {
		order := OutboundOrder{OutboundNo: fmt.Sprintf("OUT%d", time.Now().UnixNano()), WarehouseID: uint(req.WarehouseId), CustomerName: req.CustomerName, Status: int32(wms.OutboundStatus_OUTBOUND_DRAFT), OperatorUserID: req.OperatorUserId, Remark: req.Remark}
		if err := tx.Create(&order).Error; err != nil {
			return err
		}
		for _, it := range req.Items {
			item := OutboundOrderItem{OutboundOrderID: order.ID, SkuID: uint(it.SkuId), LocationID: uint(it.LocationId), LotNo: it.LotNo, ExpireDate: it.ExpireDate, Qty: it.Qty}
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
	return &wms.CreateOutboundOrderResp{Id: orderID}, nil
}

func (p *WmsProQuery) ShipOutboundOrder(req *wms.ShipOutboundOrderReq) (*wms.ShipOutboundOrderResp, error) {
	err := p.q.db.WithContext(p.q.ctx).Transaction(func(tx *gorm.DB) error {
		var order OutboundOrder
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&order, req.Id).Error; err != nil {
			return err
		}
		if order.Status == int32(wms.OutboundStatus_OUTBOUND_SHIPPED) {
			return nil
		}
		if order.Status == int32(wms.OutboundStatus_OUTBOUND_CANCELED) {
			return fmt.Errorf("outbound order canceled")
		}
		var items []OutboundOrderItem
		if err := tx.Where("outbound_order_id = ?", order.ID).Find(&items).Error; err != nil {
			return err
		}

		for _, it := range items {
			var inv Inventory
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("warehouse_id = ? and location_id = ? and sku_id = ? and lot_no = ? and expire_date = ?", order.WarehouseID, it.LocationID, it.SkuID, it.LotNo, it.ExpireDate).First(&inv).Error; err != nil {
				return fmt.Errorf("inventory not found for sku=%d location=%d", it.SkuID, it.LocationID)
			}
			if inv.AvailableQuantity < it.Qty {
				return fmt.Errorf("insufficient inventory for sku=%d location=%d", it.SkuID, it.LocationID)
			}
		}
		for _, it := range items {
			var inv Inventory
			if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("warehouse_id = ? and location_id = ? and sku_id = ? and lot_no = ? and expire_date = ?", order.WarehouseID, it.LocationID, it.SkuID, it.LotNo, it.ExpireDate).First(&inv).Error; err != nil {
				return err
			}
			before := inv.Quantity
			inv.Quantity -= it.Qty
			inv.AvailableQuantity = inv.Quantity - inv.LockedQuantity
			if err := tx.Model(&Inventory{}).Where("id = ?", inv.ID).Updates(map[string]any{"quantity": inv.Quantity, "available_quantity": inv.AvailableQuantity}).Error; err != nil {
				return err
			}
			txn := InventoryTxn{BizType: int32(wms.BizType_OUTBOUND), BizNo: order.OutboundNo, WarehouseID: order.WarehouseID, LocationID: it.LocationID, SkuID: it.SkuID, LotNo: it.LotNo, ExpireDate: it.ExpireDate, QtyChange: -it.Qty, BeforeQty: before, AfterQty: inv.Quantity, OperatorUserID: req.OperatorUserId}
			if err := tx.Create(&txn).Error; err != nil {
				return err
			}
			p.cacheDel(p.key("inventory", int64(inv.ID)))
		}
		if err := tx.Model(&OutboundOrder{}).Where("id = ?", order.ID).Updates(map[string]any{"status": int32(wms.OutboundStatus_OUTBOUND_SHIPPED), "operator_user_id": req.OperatorUserId}).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &wms.ShipOutboundOrderResp{Success: true}, nil
}

func (p *WmsProQuery) GetOutboundOrder(id int64) (*wms.GetOutboundOrderResp, error) {
	var order OutboundOrder
	if err := p.q.db.WithContext(p.q.ctx).First(&order, id).Error; err != nil {
		return nil, err
	}
	var items []OutboundOrderItem
	if err := p.q.db.WithContext(p.q.ctx).Where("outbound_order_id = ?", order.ID).Find(&items).Error; err != nil {
		return nil, err
	}
	itemList := make([]*wms.OutboundOrderItemInfo, 0, len(items))
	for _, it := range items {
		itemList = append(itemList, toOutboundItemInfo(it))
	}
	return &wms.GetOutboundOrderResp{Order: &wms.OutboundOrderInfo{Id: int64(order.ID), OutboundNo: order.OutboundNo, WarehouseId: int64(order.WarehouseID), CustomerName: order.CustomerName, Status: wms.OutboundStatus(order.Status), OperatorUserId: order.OperatorUserID, Remark: order.Remark, CreateTime: toTimeString(order.CreatedAt), UpdateTime: toTimeString(order.UpdatedAt), Items: itemList}}, nil
}

func (p *WmsProQuery) ListOutboundOrder(req *wms.ListOutboundOrderReq) (*wms.ListOutboundOrderResp, error) {
	pageNum, pageSize := pageNormalize(req.PageNum, req.PageSize)
	q := p.q.db.WithContext(p.q.ctx).Model(&OutboundOrder{})
	if req.WarehouseId > 0 {
		q = q.Where("warehouse_id = ?", req.WarehouseId)
	}
	if req.OutboundNo != "" {
		q = q.Where("outbound_no like ?", likePrefix(req.OutboundNo))
	}
	if req.Status == wms.OutboundStatus_OUTBOUND_SHIPPED || req.Status == wms.OutboundStatus_OUTBOUND_CANCELED {
		q = q.Where("status = ?", int32(req.Status))
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}
	var rows []OutboundOrder
	if err := q.Order("id desc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}
	ret := make([]*wms.OutboundOrderInfo, 0, len(rows))
	for _, row := range rows {
		var items []OutboundOrderItem
		if err := p.q.db.WithContext(p.q.ctx).Where("outbound_order_id = ?", row.ID).Find(&items).Error; err != nil {
			return nil, err
		}
		itemList := make([]*wms.OutboundOrderItemInfo, 0, len(items))
		for _, it := range items {
			itemList = append(itemList, toOutboundItemInfo(it))
		}
		ret = append(ret, &wms.OutboundOrderInfo{Id: int64(row.ID), OutboundNo: row.OutboundNo, WarehouseId: int64(row.WarehouseID), CustomerName: row.CustomerName, Status: wms.OutboundStatus(row.Status), OperatorUserId: row.OperatorUserID, Remark: row.Remark, CreateTime: toTimeString(row.CreatedAt), UpdateTime: toTimeString(row.UpdatedAt), Items: itemList})
	}
	return &wms.ListOutboundOrderResp{Records: ret, Total: total}, nil
}

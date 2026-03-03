package main

import (
	"context"
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"github.com/MoScenix/wms/app/wms/biz/service"
)

// WmsServiceImpl implements the last service interface defined in the IDL.
type WmsServiceImpl struct{}

// AddWarehouse implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) AddWarehouse(ctx context.Context, req *wms.AddWarehouseReq) (resp *wms.AddWarehouseResp, err error) {
	resp, err = service.NewAddWarehouseService(ctx).Run(req)

	return resp, err
}

// UpdateWarehouse implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) UpdateWarehouse(ctx context.Context, req *wms.UpdateWarehouseReq) (resp *wms.UpdateWarehouseResp, err error) {
	resp, err = service.NewUpdateWarehouseService(ctx).Run(req)

	return resp, err
}

// GetWarehouse implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) GetWarehouse(ctx context.Context, req *wms.GetWarehouseReq) (resp *wms.GetWarehouseResp, err error) {
	resp, err = service.NewGetWarehouseService(ctx).Run(req)

	return resp, err
}

// ListWarehouse implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) ListWarehouse(ctx context.Context, req *wms.ListWarehouseReq) (resp *wms.ListWarehouseResp, err error) {
	resp, err = service.NewListWarehouseService(ctx).Run(req)

	return resp, err
}

// AddLocation implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) AddLocation(ctx context.Context, req *wms.AddLocationReq) (resp *wms.AddLocationResp, err error) {
	resp, err = service.NewAddLocationService(ctx).Run(req)

	return resp, err
}

// UpdateLocation implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) UpdateLocation(ctx context.Context, req *wms.UpdateLocationReq) (resp *wms.UpdateLocationResp, err error) {
	resp, err = service.NewUpdateLocationService(ctx).Run(req)

	return resp, err
}

// GetLocation implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) GetLocation(ctx context.Context, req *wms.GetLocationReq) (resp *wms.GetLocationResp, err error) {
	resp, err = service.NewGetLocationService(ctx).Run(req)

	return resp, err
}

// ListLocation implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) ListLocation(ctx context.Context, req *wms.ListLocationReq) (resp *wms.ListLocationResp, err error) {
	resp, err = service.NewListLocationService(ctx).Run(req)

	return resp, err
}

// AddSku implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) AddSku(ctx context.Context, req *wms.AddSkuReq) (resp *wms.AddSkuResp, err error) {
	resp, err = service.NewAddSkuService(ctx).Run(req)

	return resp, err
}

// UpdateSku implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) UpdateSku(ctx context.Context, req *wms.UpdateSkuReq) (resp *wms.UpdateSkuResp, err error) {
	resp, err = service.NewUpdateSkuService(ctx).Run(req)

	return resp, err
}

// GetSku implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) GetSku(ctx context.Context, req *wms.GetSkuReq) (resp *wms.GetSkuResp, err error) {
	resp, err = service.NewGetSkuService(ctx).Run(req)

	return resp, err
}

// ListSku implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) ListSku(ctx context.Context, req *wms.ListSkuReq) (resp *wms.ListSkuResp, err error) {
	resp, err = service.NewListSkuService(ctx).Run(req)

	return resp, err
}

// GetInventory implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) GetInventory(ctx context.Context, req *wms.GetInventoryReq) (resp *wms.GetInventoryResp, err error) {
	resp, err = service.NewGetInventoryService(ctx).Run(req)

	return resp, err
}

// ListInventory implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) ListInventory(ctx context.Context, req *wms.ListInventoryReq) (resp *wms.ListInventoryResp, err error) {
	resp, err = service.NewListInventoryService(ctx).Run(req)

	return resp, err
}

// CheckInventory implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) CheckInventory(ctx context.Context, req *wms.CheckInventoryReq) (resp *wms.CheckInventoryResp, err error) {
	resp, err = service.NewCheckInventoryService(ctx).Run(req)

	return resp, err
}

// CreateInboundOrder implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) CreateInboundOrder(ctx context.Context, req *wms.CreateInboundOrderReq) (resp *wms.CreateInboundOrderResp, err error) {
	resp, err = service.NewCreateInboundOrderService(ctx).Run(req)

	return resp, err
}

// ReceiveInboundOrder implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) ReceiveInboundOrder(ctx context.Context, req *wms.ReceiveInboundOrderReq) (resp *wms.ReceiveInboundOrderResp, err error) {
	resp, err = service.NewReceiveInboundOrderService(ctx).Run(req)

	return resp, err
}

// GetInboundOrder implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) GetInboundOrder(ctx context.Context, req *wms.GetInboundOrderReq) (resp *wms.GetInboundOrderResp, err error) {
	resp, err = service.NewGetInboundOrderService(ctx).Run(req)

	return resp, err
}

// ListInboundOrder implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) ListInboundOrder(ctx context.Context, req *wms.ListInboundOrderReq) (resp *wms.ListInboundOrderResp, err error) {
	resp, err = service.NewListInboundOrderService(ctx).Run(req)

	return resp, err
}

// CreateOutboundOrder implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) CreateOutboundOrder(ctx context.Context, req *wms.CreateOutboundOrderReq) (resp *wms.CreateOutboundOrderResp, err error) {
	resp, err = service.NewCreateOutboundOrderService(ctx).Run(req)

	return resp, err
}

// ShipOutboundOrder implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) ShipOutboundOrder(ctx context.Context, req *wms.ShipOutboundOrderReq) (resp *wms.ShipOutboundOrderResp, err error) {
	resp, err = service.NewShipOutboundOrderService(ctx).Run(req)

	return resp, err
}

// GetOutboundOrder implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) GetOutboundOrder(ctx context.Context, req *wms.GetOutboundOrderReq) (resp *wms.GetOutboundOrderResp, err error) {
	resp, err = service.NewGetOutboundOrderService(ctx).Run(req)

	return resp, err
}

// ListOutboundOrder implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) ListOutboundOrder(ctx context.Context, req *wms.ListOutboundOrderReq) (resp *wms.ListOutboundOrderResp, err error) {
	resp, err = service.NewListOutboundOrderService(ctx).Run(req)

	return resp, err
}

// ListInventoryTxn implements the WmsServiceImpl interface.
func (s *WmsServiceImpl) ListInventoryTxn(ctx context.Context, req *wms.ListInventoryTxnReq) (resp *wms.ListInventoryTxnResp, err error) {
	resp, err = service.NewListInventoryTxnService(ctx).Run(req)

	return resp, err
}

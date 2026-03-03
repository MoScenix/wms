package wms

import (
	"context"
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"

	"github.com/MoScenix/wms/rpc_gen/kitex_gen/wms/wmsservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() wmsservice.Client
	Service() string
	AddWarehouse(ctx context.Context, Req *wms.AddWarehouseReq, callOptions ...callopt.Option) (r *wms.AddWarehouseResp, err error)
	UpdateWarehouse(ctx context.Context, Req *wms.UpdateWarehouseReq, callOptions ...callopt.Option) (r *wms.UpdateWarehouseResp, err error)
	GetWarehouse(ctx context.Context, Req *wms.GetWarehouseReq, callOptions ...callopt.Option) (r *wms.GetWarehouseResp, err error)
	ListWarehouse(ctx context.Context, Req *wms.ListWarehouseReq, callOptions ...callopt.Option) (r *wms.ListWarehouseResp, err error)
	AddLocation(ctx context.Context, Req *wms.AddLocationReq, callOptions ...callopt.Option) (r *wms.AddLocationResp, err error)
	UpdateLocation(ctx context.Context, Req *wms.UpdateLocationReq, callOptions ...callopt.Option) (r *wms.UpdateLocationResp, err error)
	GetLocation(ctx context.Context, Req *wms.GetLocationReq, callOptions ...callopt.Option) (r *wms.GetLocationResp, err error)
	ListLocation(ctx context.Context, Req *wms.ListLocationReq, callOptions ...callopt.Option) (r *wms.ListLocationResp, err error)
	AddSku(ctx context.Context, Req *wms.AddSkuReq, callOptions ...callopt.Option) (r *wms.AddSkuResp, err error)
	UpdateSku(ctx context.Context, Req *wms.UpdateSkuReq, callOptions ...callopt.Option) (r *wms.UpdateSkuResp, err error)
	GetSku(ctx context.Context, Req *wms.GetSkuReq, callOptions ...callopt.Option) (r *wms.GetSkuResp, err error)
	ListSku(ctx context.Context, Req *wms.ListSkuReq, callOptions ...callopt.Option) (r *wms.ListSkuResp, err error)
	GetInventory(ctx context.Context, Req *wms.GetInventoryReq, callOptions ...callopt.Option) (r *wms.GetInventoryResp, err error)
	ListInventory(ctx context.Context, Req *wms.ListInventoryReq, callOptions ...callopt.Option) (r *wms.ListInventoryResp, err error)
	CheckInventory(ctx context.Context, Req *wms.CheckInventoryReq, callOptions ...callopt.Option) (r *wms.CheckInventoryResp, err error)
	CreateInboundOrder(ctx context.Context, Req *wms.CreateInboundOrderReq, callOptions ...callopt.Option) (r *wms.CreateInboundOrderResp, err error)
	ReceiveInboundOrder(ctx context.Context, Req *wms.ReceiveInboundOrderReq, callOptions ...callopt.Option) (r *wms.ReceiveInboundOrderResp, err error)
	GetInboundOrder(ctx context.Context, Req *wms.GetInboundOrderReq, callOptions ...callopt.Option) (r *wms.GetInboundOrderResp, err error)
	ListInboundOrder(ctx context.Context, Req *wms.ListInboundOrderReq, callOptions ...callopt.Option) (r *wms.ListInboundOrderResp, err error)
	CreateOutboundOrder(ctx context.Context, Req *wms.CreateOutboundOrderReq, callOptions ...callopt.Option) (r *wms.CreateOutboundOrderResp, err error)
	ShipOutboundOrder(ctx context.Context, Req *wms.ShipOutboundOrderReq, callOptions ...callopt.Option) (r *wms.ShipOutboundOrderResp, err error)
	GetOutboundOrder(ctx context.Context, Req *wms.GetOutboundOrderReq, callOptions ...callopt.Option) (r *wms.GetOutboundOrderResp, err error)
	ListOutboundOrder(ctx context.Context, Req *wms.ListOutboundOrderReq, callOptions ...callopt.Option) (r *wms.ListOutboundOrderResp, err error)
	ListInventoryTxn(ctx context.Context, Req *wms.ListInventoryTxnReq, callOptions ...callopt.Option) (r *wms.ListInventoryTxnResp, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := wmsservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient wmsservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() wmsservice.Client {
	return c.kitexClient
}

func (c *clientImpl) AddWarehouse(ctx context.Context, Req *wms.AddWarehouseReq, callOptions ...callopt.Option) (r *wms.AddWarehouseResp, err error) {
	return c.kitexClient.AddWarehouse(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateWarehouse(ctx context.Context, Req *wms.UpdateWarehouseReq, callOptions ...callopt.Option) (r *wms.UpdateWarehouseResp, err error) {
	return c.kitexClient.UpdateWarehouse(ctx, Req, callOptions...)
}

func (c *clientImpl) GetWarehouse(ctx context.Context, Req *wms.GetWarehouseReq, callOptions ...callopt.Option) (r *wms.GetWarehouseResp, err error) {
	return c.kitexClient.GetWarehouse(ctx, Req, callOptions...)
}

func (c *clientImpl) ListWarehouse(ctx context.Context, Req *wms.ListWarehouseReq, callOptions ...callopt.Option) (r *wms.ListWarehouseResp, err error) {
	return c.kitexClient.ListWarehouse(ctx, Req, callOptions...)
}

func (c *clientImpl) AddLocation(ctx context.Context, Req *wms.AddLocationReq, callOptions ...callopt.Option) (r *wms.AddLocationResp, err error) {
	return c.kitexClient.AddLocation(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateLocation(ctx context.Context, Req *wms.UpdateLocationReq, callOptions ...callopt.Option) (r *wms.UpdateLocationResp, err error) {
	return c.kitexClient.UpdateLocation(ctx, Req, callOptions...)
}

func (c *clientImpl) GetLocation(ctx context.Context, Req *wms.GetLocationReq, callOptions ...callopt.Option) (r *wms.GetLocationResp, err error) {
	return c.kitexClient.GetLocation(ctx, Req, callOptions...)
}

func (c *clientImpl) ListLocation(ctx context.Context, Req *wms.ListLocationReq, callOptions ...callopt.Option) (r *wms.ListLocationResp, err error) {
	return c.kitexClient.ListLocation(ctx, Req, callOptions...)
}

func (c *clientImpl) AddSku(ctx context.Context, Req *wms.AddSkuReq, callOptions ...callopt.Option) (r *wms.AddSkuResp, err error) {
	return c.kitexClient.AddSku(ctx, Req, callOptions...)
}

func (c *clientImpl) UpdateSku(ctx context.Context, Req *wms.UpdateSkuReq, callOptions ...callopt.Option) (r *wms.UpdateSkuResp, err error) {
	return c.kitexClient.UpdateSku(ctx, Req, callOptions...)
}

func (c *clientImpl) GetSku(ctx context.Context, Req *wms.GetSkuReq, callOptions ...callopt.Option) (r *wms.GetSkuResp, err error) {
	return c.kitexClient.GetSku(ctx, Req, callOptions...)
}

func (c *clientImpl) ListSku(ctx context.Context, Req *wms.ListSkuReq, callOptions ...callopt.Option) (r *wms.ListSkuResp, err error) {
	return c.kitexClient.ListSku(ctx, Req, callOptions...)
}

func (c *clientImpl) GetInventory(ctx context.Context, Req *wms.GetInventoryReq, callOptions ...callopt.Option) (r *wms.GetInventoryResp, err error) {
	return c.kitexClient.GetInventory(ctx, Req, callOptions...)
}

func (c *clientImpl) ListInventory(ctx context.Context, Req *wms.ListInventoryReq, callOptions ...callopt.Option) (r *wms.ListInventoryResp, err error) {
	return c.kitexClient.ListInventory(ctx, Req, callOptions...)
}

func (c *clientImpl) CheckInventory(ctx context.Context, Req *wms.CheckInventoryReq, callOptions ...callopt.Option) (r *wms.CheckInventoryResp, err error) {
	return c.kitexClient.CheckInventory(ctx, Req, callOptions...)
}

func (c *clientImpl) CreateInboundOrder(ctx context.Context, Req *wms.CreateInboundOrderReq, callOptions ...callopt.Option) (r *wms.CreateInboundOrderResp, err error) {
	return c.kitexClient.CreateInboundOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) ReceiveInboundOrder(ctx context.Context, Req *wms.ReceiveInboundOrderReq, callOptions ...callopt.Option) (r *wms.ReceiveInboundOrderResp, err error) {
	return c.kitexClient.ReceiveInboundOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) GetInboundOrder(ctx context.Context, Req *wms.GetInboundOrderReq, callOptions ...callopt.Option) (r *wms.GetInboundOrderResp, err error) {
	return c.kitexClient.GetInboundOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) ListInboundOrder(ctx context.Context, Req *wms.ListInboundOrderReq, callOptions ...callopt.Option) (r *wms.ListInboundOrderResp, err error) {
	return c.kitexClient.ListInboundOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) CreateOutboundOrder(ctx context.Context, Req *wms.CreateOutboundOrderReq, callOptions ...callopt.Option) (r *wms.CreateOutboundOrderResp, err error) {
	return c.kitexClient.CreateOutboundOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) ShipOutboundOrder(ctx context.Context, Req *wms.ShipOutboundOrderReq, callOptions ...callopt.Option) (r *wms.ShipOutboundOrderResp, err error) {
	return c.kitexClient.ShipOutboundOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) GetOutboundOrder(ctx context.Context, Req *wms.GetOutboundOrderReq, callOptions ...callopt.Option) (r *wms.GetOutboundOrderResp, err error) {
	return c.kitexClient.GetOutboundOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) ListOutboundOrder(ctx context.Context, Req *wms.ListOutboundOrderReq, callOptions ...callopt.Option) (r *wms.ListOutboundOrderResp, err error) {
	return c.kitexClient.ListOutboundOrder(ctx, Req, callOptions...)
}

func (c *clientImpl) ListInventoryTxn(ctx context.Context, Req *wms.ListInventoryTxnReq, callOptions ...callopt.Option) (r *wms.ListInventoryTxnResp, err error) {
	return c.kitexClient.ListInventoryTxn(ctx, Req, callOptions...)
}

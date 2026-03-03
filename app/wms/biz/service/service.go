package service

import (
	"context"

	"github.com/MoScenix/wms/app/wms/biz/dal/mysql"
	"github.com/MoScenix/wms/app/wms/biz/dal/redis"
	"github.com/MoScenix/wms/app/wms/biz/model"
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
)

type AddWarehouseService struct{ ctx context.Context }
type UpdateWarehouseService struct{ ctx context.Context }
type GetWarehouseService struct{ ctx context.Context }
type ListWarehouseService struct{ ctx context.Context }
type AddLocationService struct{ ctx context.Context }
type UpdateLocationService struct{ ctx context.Context }
type GetLocationService struct{ ctx context.Context }
type ListLocationService struct{ ctx context.Context }
type AddSkuService struct{ ctx context.Context }
type UpdateSkuService struct{ ctx context.Context }
type GetSkuService struct{ ctx context.Context }
type ListSkuService struct{ ctx context.Context }
type GetInventoryService struct{ ctx context.Context }
type ListInventoryService struct{ ctx context.Context }
type CheckInventoryService struct{ ctx context.Context }
type CreateInboundOrderService struct{ ctx context.Context }
type ReceiveInboundOrderService struct{ ctx context.Context }
type GetInboundOrderService struct{ ctx context.Context }
type ListInboundOrderService struct{ ctx context.Context }
type CreateOutboundOrderService struct{ ctx context.Context }
type ShipOutboundOrderService struct{ ctx context.Context }
type GetOutboundOrderService struct{ ctx context.Context }
type ListOutboundOrderService struct{ ctx context.Context }
type ListInventoryTxnService struct{ ctx context.Context }

func NewAddWarehouseService(ctx context.Context) *AddWarehouseService {
	return &AddWarehouseService{ctx: ctx}
}
func NewUpdateWarehouseService(ctx context.Context) *UpdateWarehouseService {
	return &UpdateWarehouseService{ctx: ctx}
}
func NewGetWarehouseService(ctx context.Context) *GetWarehouseService {
	return &GetWarehouseService{ctx: ctx}
}
func NewListWarehouseService(ctx context.Context) *ListWarehouseService {
	return &ListWarehouseService{ctx: ctx}
}
func NewAddLocationService(ctx context.Context) *AddLocationService {
	return &AddLocationService{ctx: ctx}
}
func NewUpdateLocationService(ctx context.Context) *UpdateLocationService {
	return &UpdateLocationService{ctx: ctx}
}
func NewGetLocationService(ctx context.Context) *GetLocationService {
	return &GetLocationService{ctx: ctx}
}
func NewListLocationService(ctx context.Context) *ListLocationService {
	return &ListLocationService{ctx: ctx}
}
func NewAddSkuService(ctx context.Context) *AddSkuService       { return &AddSkuService{ctx: ctx} }
func NewUpdateSkuService(ctx context.Context) *UpdateSkuService { return &UpdateSkuService{ctx: ctx} }
func NewGetSkuService(ctx context.Context) *GetSkuService       { return &GetSkuService{ctx: ctx} }
func NewListSkuService(ctx context.Context) *ListSkuService     { return &ListSkuService{ctx: ctx} }
func NewGetInventoryService(ctx context.Context) *GetInventoryService {
	return &GetInventoryService{ctx: ctx}
}
func NewListInventoryService(ctx context.Context) *ListInventoryService {
	return &ListInventoryService{ctx: ctx}
}
func NewCheckInventoryService(ctx context.Context) *CheckInventoryService {
	return &CheckInventoryService{ctx: ctx}
}
func NewCreateInboundOrderService(ctx context.Context) *CreateInboundOrderService {
	return &CreateInboundOrderService{ctx: ctx}
}
func NewReceiveInboundOrderService(ctx context.Context) *ReceiveInboundOrderService {
	return &ReceiveInboundOrderService{ctx: ctx}
}
func NewGetInboundOrderService(ctx context.Context) *GetInboundOrderService {
	return &GetInboundOrderService{ctx: ctx}
}
func NewListInboundOrderService(ctx context.Context) *ListInboundOrderService {
	return &ListInboundOrderService{ctx: ctx}
}
func NewCreateOutboundOrderService(ctx context.Context) *CreateOutboundOrderService {
	return &CreateOutboundOrderService{ctx: ctx}
}
func NewShipOutboundOrderService(ctx context.Context) *ShipOutboundOrderService {
	return &ShipOutboundOrderService{ctx: ctx}
}
func NewGetOutboundOrderService(ctx context.Context) *GetOutboundOrderService {
	return &GetOutboundOrderService{ctx: ctx}
}
func NewListOutboundOrderService(ctx context.Context) *ListOutboundOrderService {
	return &ListOutboundOrderService{ctx: ctx}
}
func NewListInventoryTxnService(ctx context.Context) *ListInventoryTxnService {
	return &ListInventoryTxnService{ctx: ctx}
}

func pro(ctx context.Context) *model.WmsProQuery {
	return model.NewWmsProQuery(ctx, mysql.DB, redis.RedisClient)
}

func (s *AddWarehouseService) Run(req *wms.AddWarehouseReq) (*wms.AddWarehouseResp, error) {
	return pro(s.ctx).AddWarehouse(req)
}
func (s *UpdateWarehouseService) Run(req *wms.UpdateWarehouseReq) (*wms.UpdateWarehouseResp, error) {
	return pro(s.ctx).UpdateWarehouse(req)
}
func (s *GetWarehouseService) Run(req *wms.GetWarehouseReq) (*wms.GetWarehouseResp, error) {
	return pro(s.ctx).GetWarehouse(req.GetId())
}
func (s *ListWarehouseService) Run(req *wms.ListWarehouseReq) (*wms.ListWarehouseResp, error) {
	return pro(s.ctx).ListWarehouse(req)
}
func (s *AddLocationService) Run(req *wms.AddLocationReq) (*wms.AddLocationResp, error) {
	return pro(s.ctx).AddLocation(req)
}
func (s *UpdateLocationService) Run(req *wms.UpdateLocationReq) (*wms.UpdateLocationResp, error) {
	return pro(s.ctx).UpdateLocation(req)
}
func (s *GetLocationService) Run(req *wms.GetLocationReq) (*wms.GetLocationResp, error) {
	return pro(s.ctx).GetLocation(req.GetId())
}
func (s *ListLocationService) Run(req *wms.ListLocationReq) (*wms.ListLocationResp, error) {
	return pro(s.ctx).ListLocation(req)
}
func (s *AddSkuService) Run(req *wms.AddSkuReq) (*wms.AddSkuResp, error) {
	return pro(s.ctx).AddSku(req)
}
func (s *UpdateSkuService) Run(req *wms.UpdateSkuReq) (*wms.UpdateSkuResp, error) {
	return pro(s.ctx).UpdateSku(req)
}
func (s *GetSkuService) Run(req *wms.GetSkuReq) (*wms.GetSkuResp, error) {
	return pro(s.ctx).GetSku(req.GetId())
}
func (s *ListSkuService) Run(req *wms.ListSkuReq) (*wms.ListSkuResp, error) {
	return pro(s.ctx).ListSku(req)
}
func (s *GetInventoryService) Run(req *wms.GetInventoryReq) (*wms.GetInventoryResp, error) {
	return pro(s.ctx).GetInventory(req.GetId())
}
func (s *ListInventoryService) Run(req *wms.ListInventoryReq) (*wms.ListInventoryResp, error) {
	return pro(s.ctx).ListInventory(req)
}
func (s *CheckInventoryService) Run(req *wms.CheckInventoryReq) (*wms.CheckInventoryResp, error) {
	return pro(s.ctx).CheckInventory(req)
}
func (s *CreateInboundOrderService) Run(req *wms.CreateInboundOrderReq) (*wms.CreateInboundOrderResp, error) {
	return pro(s.ctx).CreateInboundOrder(req)
}
func (s *ReceiveInboundOrderService) Run(req *wms.ReceiveInboundOrderReq) (*wms.ReceiveInboundOrderResp, error) {
	return pro(s.ctx).ReceiveInboundOrder(req)
}
func (s *GetInboundOrderService) Run(req *wms.GetInboundOrderReq) (*wms.GetInboundOrderResp, error) {
	return pro(s.ctx).GetInboundOrder(req.GetId())
}
func (s *ListInboundOrderService) Run(req *wms.ListInboundOrderReq) (*wms.ListInboundOrderResp, error) {
	return pro(s.ctx).ListInboundOrder(req)
}
func (s *CreateOutboundOrderService) Run(req *wms.CreateOutboundOrderReq) (*wms.CreateOutboundOrderResp, error) {
	return pro(s.ctx).CreateOutboundOrder(req)
}
func (s *ShipOutboundOrderService) Run(req *wms.ShipOutboundOrderReq) (*wms.ShipOutboundOrderResp, error) {
	return pro(s.ctx).ShipOutboundOrder(req)
}
func (s *GetOutboundOrderService) Run(req *wms.GetOutboundOrderReq) (*wms.GetOutboundOrderResp, error) {
	return pro(s.ctx).GetOutboundOrder(req.GetId())
}
func (s *ListOutboundOrderService) Run(req *wms.ListOutboundOrderReq) (*wms.ListOutboundOrderResp, error) {
	return pro(s.ctx).ListOutboundOrder(req)
}
func (s *ListInventoryTxnService) Run(req *wms.ListInventoryTxnReq) (*wms.ListInventoryTxnResp, error) {
	return pro(s.ctx).ListInventoryTxn(req)
}

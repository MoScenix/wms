package service

import (
	"context"
	"math"

	"github.com/MoScenix/wms/app/bff/biz/utils"
	bwms "github.com/MoScenix/wms/app/bff/hertz_gen/bff/wms"
	"github.com/MoScenix/wms/app/bff/infra/rpc"
	rpcwms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"github.com/cloudwego/hertz/pkg/app"
)

type AddWarehouseService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type UpdateWarehouseService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type GetWarehouseByIdService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type ListWarehouseByPageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type AddLocationService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type UpdateLocationService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type GetLocationByIdService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type ListLocationByPageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type AddSkuService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type UpdateSkuService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type GetSkuByIdService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type ListSkuByPageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type GetInventoryByIdService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type ListInventoryByPageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type CheckInventoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type CreateInboundOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type ReceiveInboundOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type GetInboundOrderByIdService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type ListInboundOrderByPageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type CreateOutboundOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type ShipOutboundOrderService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type GetOutboundOrderByIdService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type ListOutboundOrderByPageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}
type ListInventoryTxnByPageService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewAddWarehouseService(ctx context.Context, reqCtx *app.RequestContext) *AddWarehouseService {
	return &AddWarehouseService{RequestContext: reqCtx, Context: ctx}
}
func NewUpdateWarehouseService(ctx context.Context, reqCtx *app.RequestContext) *UpdateWarehouseService {
	return &UpdateWarehouseService{RequestContext: reqCtx, Context: ctx}
}
func NewGetWarehouseByIdService(ctx context.Context, reqCtx *app.RequestContext) *GetWarehouseByIdService {
	return &GetWarehouseByIdService{RequestContext: reqCtx, Context: ctx}
}
func NewListWarehouseByPageService(ctx context.Context, reqCtx *app.RequestContext) *ListWarehouseByPageService {
	return &ListWarehouseByPageService{RequestContext: reqCtx, Context: ctx}
}
func NewAddLocationService(ctx context.Context, reqCtx *app.RequestContext) *AddLocationService {
	return &AddLocationService{RequestContext: reqCtx, Context: ctx}
}
func NewUpdateLocationService(ctx context.Context, reqCtx *app.RequestContext) *UpdateLocationService {
	return &UpdateLocationService{RequestContext: reqCtx, Context: ctx}
}
func NewGetLocationByIdService(ctx context.Context, reqCtx *app.RequestContext) *GetLocationByIdService {
	return &GetLocationByIdService{RequestContext: reqCtx, Context: ctx}
}
func NewListLocationByPageService(ctx context.Context, reqCtx *app.RequestContext) *ListLocationByPageService {
	return &ListLocationByPageService{RequestContext: reqCtx, Context: ctx}
}
func NewAddSkuService(ctx context.Context, reqCtx *app.RequestContext) *AddSkuService {
	return &AddSkuService{RequestContext: reqCtx, Context: ctx}
}
func NewUpdateSkuService(ctx context.Context, reqCtx *app.RequestContext) *UpdateSkuService {
	return &UpdateSkuService{RequestContext: reqCtx, Context: ctx}
}
func NewGetSkuByIdService(ctx context.Context, reqCtx *app.RequestContext) *GetSkuByIdService {
	return &GetSkuByIdService{RequestContext: reqCtx, Context: ctx}
}
func NewListSkuByPageService(ctx context.Context, reqCtx *app.RequestContext) *ListSkuByPageService {
	return &ListSkuByPageService{RequestContext: reqCtx, Context: ctx}
}
func NewGetInventoryByIdService(ctx context.Context, reqCtx *app.RequestContext) *GetInventoryByIdService {
	return &GetInventoryByIdService{RequestContext: reqCtx, Context: ctx}
}
func NewListInventoryByPageService(ctx context.Context, reqCtx *app.RequestContext) *ListInventoryByPageService {
	return &ListInventoryByPageService{RequestContext: reqCtx, Context: ctx}
}
func NewCheckInventoryService(ctx context.Context, reqCtx *app.RequestContext) *CheckInventoryService {
	return &CheckInventoryService{RequestContext: reqCtx, Context: ctx}
}
func NewCreateInboundOrderService(ctx context.Context, reqCtx *app.RequestContext) *CreateInboundOrderService {
	return &CreateInboundOrderService{RequestContext: reqCtx, Context: ctx}
}
func NewReceiveInboundOrderService(ctx context.Context, reqCtx *app.RequestContext) *ReceiveInboundOrderService {
	return &ReceiveInboundOrderService{RequestContext: reqCtx, Context: ctx}
}
func NewGetInboundOrderByIdService(ctx context.Context, reqCtx *app.RequestContext) *GetInboundOrderByIdService {
	return &GetInboundOrderByIdService{RequestContext: reqCtx, Context: ctx}
}
func NewListInboundOrderByPageService(ctx context.Context, reqCtx *app.RequestContext) *ListInboundOrderByPageService {
	return &ListInboundOrderByPageService{RequestContext: reqCtx, Context: ctx}
}
func NewCreateOutboundOrderService(ctx context.Context, reqCtx *app.RequestContext) *CreateOutboundOrderService {
	return &CreateOutboundOrderService{RequestContext: reqCtx, Context: ctx}
}
func NewShipOutboundOrderService(ctx context.Context, reqCtx *app.RequestContext) *ShipOutboundOrderService {
	return &ShipOutboundOrderService{RequestContext: reqCtx, Context: ctx}
}
func NewGetOutboundOrderByIdService(ctx context.Context, reqCtx *app.RequestContext) *GetOutboundOrderByIdService {
	return &GetOutboundOrderByIdService{RequestContext: reqCtx, Context: ctx}
}
func NewListOutboundOrderByPageService(ctx context.Context, reqCtx *app.RequestContext) *ListOutboundOrderByPageService {
	return &ListOutboundOrderByPageService{RequestContext: reqCtx, Context: ctx}
}
func NewListInventoryTxnByPageService(ctx context.Context, reqCtx *app.RequestContext) *ListInventoryTxnByPageService {
	return &ListInventoryTxnByPageService{RequestContext: reqCtx, Context: ctx}
}

func codeOK() int32 { return 0 }

func getUserID(ctx context.Context) int64 {
	v := ctx.Value(utils.UserIdKey)
	switch value := v.(type) {
	case int64:
		return value
	case int32:
		return int64(value)
	case int:
		return int64(value)
	case uint64:
		return int64(value)
	case uint32:
		return int64(value)
	case uint:
		return int64(value)
	case float64:
		return int64(value)
	case float32:
		return int64(value)
	default:
		return 0
	}
}

func toWarehouse(in *rpcwms.WarehouseInfo) *bwms.Warehouse {
	if in == nil {
		return nil
	}
	return &bwms.Warehouse{Id: in.Id, WarehouseCode: in.WarehouseCode, WarehouseName: in.WarehouseName, Address: in.Address, ManagerUserId: in.ManagerUserId, CreateTime: in.CreateTime, UpdateTime: in.UpdateTime}
}
func toLocation(in *rpcwms.LocationInfo) *bwms.Location {
	if in == nil {
		return nil
	}
	return &bwms.Location{Id: in.Id, WarehouseId: in.WarehouseId, LocationCode: in.LocationCode, LocationName: in.LocationName, LocationType: in.LocationType, CreateTime: in.CreateTime, UpdateTime: in.UpdateTime}
}
func toSku(in *rpcwms.SkuInfo) *bwms.Sku {
	if in == nil {
		return nil
	}
	return &bwms.Sku{Id: in.Id, SkuCode: in.SkuCode, SkuName: in.SkuName, Unit: in.Unit, Category: in.Category, CreateTime: in.CreateTime, UpdateTime: in.UpdateTime}
}
func toInventory(in *rpcwms.InventoryInfo) *bwms.Inventory {
	if in == nil {
		return nil
	}
	return &bwms.Inventory{Id: in.Id, WarehouseId: in.WarehouseId, LocationId: in.LocationId, SkuId: in.SkuId, LotNo: in.LotNo, ExpireDate: in.ExpireDate, Quantity: in.Quantity, LockedQuantity: in.LockedQuantity, AvailableQuantity: in.AvailableQuantity, CreateTime: in.CreateTime, UpdateTime: in.UpdateTime}
}
func toInboundOrder(in *rpcwms.InboundOrderInfo) *bwms.InboundOrder {
	if in == nil {
		return nil
	}
	items := make([]*bwms.InboundOrderItem, 0, len(in.Items))
	for _, item := range in.Items {
		items = append(items, &bwms.InboundOrderItem{Id: item.Id, InboundOrderId: item.InboundOrderId, SkuId: item.SkuId, LocationId: item.LocationId, LotNo: item.LotNo, ExpireDate: item.ExpireDate, Qty: item.Qty})
	}
	return &bwms.InboundOrder{Id: in.Id, InboundNo: in.InboundNo, WarehouseId: in.WarehouseId, SupplierName: in.SupplierName, Status: int32(in.Status), OperatorUserId: in.OperatorUserId, Remark: in.Remark, CreateTime: in.CreateTime, UpdateTime: in.UpdateTime, Items: items}
}
func toOutboundOrder(in *rpcwms.OutboundOrderInfo) *bwms.OutboundOrder {
	if in == nil {
		return nil
	}
	items := make([]*bwms.OutboundOrderItem, 0, len(in.Items))
	for _, item := range in.Items {
		items = append(items, &bwms.OutboundOrderItem{Id: item.Id, OutboundOrderId: item.OutboundOrderId, SkuId: item.SkuId, LocationId: item.LocationId, LotNo: item.LotNo, ExpireDate: item.ExpireDate, Qty: item.Qty})
	}
	return &bwms.OutboundOrder{Id: in.Id, OutboundNo: in.OutboundNo, WarehouseId: in.WarehouseId, CustomerName: in.CustomerName, Status: int32(in.Status), OperatorUserId: in.OperatorUserId, Remark: in.Remark, CreateTime: in.CreateTime, UpdateTime: in.UpdateTime, Items: items}
}

func toTotalPage(total, pageSize int64) int64 {
	if pageSize <= 0 {
		pageSize = 10
	}
	return int64(math.Ceil(float64(total) / float64(pageSize)))
}

func (h *AddWarehouseService) Run(req *bwms.WarehouseAddRequest) (*bwms.BaseResponseLong, error) {
	resp, err := rpc.WmsClient.AddWarehouse(h.Context, &rpcwms.AddWarehouseReq{WarehouseCode: req.WarehouseCode, WarehouseName: req.WarehouseName, Address: req.Address, ManagerUserId: req.ManagerUserId})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseLong{Code: codeOK(), Data: resp.Id, Message: "ok"}, nil
}
func (h *UpdateWarehouseService) Run(req *bwms.WarehouseUpdateRequest) (*bwms.BaseResponseBoolean, error) {
	resp, err := rpc.WmsClient.UpdateWarehouse(h.Context, &rpcwms.UpdateWarehouseReq{Id: req.Id, WarehouseName: req.WarehouseName, Address: req.Address, ManagerUserId: req.ManagerUserId})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseBoolean{Code: codeOK(), Data: resp.Success, Message: "ok"}, nil
}
func (h *GetWarehouseByIdService) Run(req *bwms.GetByIdRequest) (*bwms.BaseResponseWarehouse, error) {
	resp, err := rpc.WmsClient.GetWarehouse(h.Context, &rpcwms.GetWarehouseReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseWarehouse{Code: codeOK(), Data: toWarehouse(resp.Warehouse), Message: "ok"}, nil
}
func (h *ListWarehouseByPageService) Run(req *bwms.WarehouseQueryRequest) (*bwms.BaseResponsePageWarehouse, error) {
	resp, err := rpc.WmsClient.ListWarehouse(h.Context, &rpcwms.ListWarehouseReq{PageNum: req.PageNum, PageSize: req.PageSize, WarehouseCode: req.WarehouseCode, WarehouseName: req.WarehouseName})
	if err != nil {
		return nil, err
	}
	records := make([]*bwms.Warehouse, 0, len(resp.Records))
	for _, rec := range resp.Records {
		records = append(records, toWarehouse(rec))
	}
	return &bwms.BaseResponsePageWarehouse{Code: codeOK(), Data: &bwms.PageWarehouse{Records: records, PageNumber: req.PageNum, PageSize: req.PageSize, TotalPage: toTotalPage(resp.Total, req.PageSize), TotalRow: resp.Total}, Message: "ok"}, nil
}

func (h *AddLocationService) Run(req *bwms.LocationAddRequest) (*bwms.BaseResponseLong, error) {
	resp, err := rpc.WmsClient.AddLocation(h.Context, &rpcwms.AddLocationReq{WarehouseId: req.WarehouseId, LocationCode: req.LocationCode, LocationName: req.LocationName, LocationType: req.LocationType})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseLong{Code: codeOK(), Data: resp.Id, Message: "ok"}, nil
}
func (h *UpdateLocationService) Run(req *bwms.LocationUpdateRequest) (*bwms.BaseResponseBoolean, error) {
	resp, err := rpc.WmsClient.UpdateLocation(h.Context, &rpcwms.UpdateLocationReq{Id: req.Id, LocationName: req.LocationName, LocationType: req.LocationType})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseBoolean{Code: codeOK(), Data: resp.Success, Message: "ok"}, nil
}
func (h *GetLocationByIdService) Run(req *bwms.GetByIdRequest) (*bwms.BaseResponseLocation, error) {
	resp, err := rpc.WmsClient.GetLocation(h.Context, &rpcwms.GetLocationReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseLocation{Code: codeOK(), Data: toLocation(resp.Location), Message: "ok"}, nil
}
func (h *ListLocationByPageService) Run(req *bwms.LocationQueryRequest) (*bwms.BaseResponsePageLocation, error) {
	resp, err := rpc.WmsClient.ListLocation(h.Context, &rpcwms.ListLocationReq{PageNum: req.PageNum, PageSize: req.PageSize, WarehouseId: req.WarehouseId, LocationCode: req.LocationCode, LocationName: req.LocationName})
	if err != nil {
		return nil, err
	}
	records := make([]*bwms.Location, 0, len(resp.Records))
	for _, rec := range resp.Records {
		records = append(records, toLocation(rec))
	}
	return &bwms.BaseResponsePageLocation{Code: codeOK(), Data: &bwms.PageLocation{Records: records, PageNumber: req.PageNum, PageSize: req.PageSize, TotalPage: toTotalPage(resp.Total, req.PageSize), TotalRow: resp.Total}, Message: "ok"}, nil
}

func (h *AddSkuService) Run(req *bwms.SkuAddRequest) (*bwms.BaseResponseLong, error) {
	resp, err := rpc.WmsClient.AddSku(h.Context, &rpcwms.AddSkuReq{SkuCode: req.SkuCode, SkuName: req.SkuName, Unit: req.Unit, Category: req.Category})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseLong{Code: codeOK(), Data: resp.Id, Message: "ok"}, nil
}
func (h *UpdateSkuService) Run(req *bwms.SkuUpdateRequest) (*bwms.BaseResponseBoolean, error) {
	resp, err := rpc.WmsClient.UpdateSku(h.Context, &rpcwms.UpdateSkuReq{Id: req.Id, SkuName: req.SkuName, Unit: req.Unit, Category: req.Category})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseBoolean{Code: codeOK(), Data: resp.Success, Message: "ok"}, nil
}
func (h *GetSkuByIdService) Run(req *bwms.GetByIdRequest) (*bwms.BaseResponseSku, error) {
	resp, err := rpc.WmsClient.GetSku(h.Context, &rpcwms.GetSkuReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseSku{Code: codeOK(), Data: toSku(resp.Sku), Message: "ok"}, nil
}
func (h *ListSkuByPageService) Run(req *bwms.SkuQueryRequest) (*bwms.BaseResponsePageSku, error) {
	resp, err := rpc.WmsClient.ListSku(h.Context, &rpcwms.ListSkuReq{PageNum: req.PageNum, PageSize: req.PageSize, SkuCode: req.SkuCode, SkuName: req.SkuName})
	if err != nil {
		return nil, err
	}
	records := make([]*bwms.Sku, 0, len(resp.Records))
	for _, rec := range resp.Records {
		records = append(records, toSku(rec))
	}
	return &bwms.BaseResponsePageSku{Code: codeOK(), Data: &bwms.PageSku{Records: records, PageNumber: req.PageNum, PageSize: req.PageSize, TotalPage: toTotalPage(resp.Total, req.PageSize), TotalRow: resp.Total}, Message: "ok"}, nil
}

func (h *GetInventoryByIdService) Run(req *bwms.GetByIdRequest) (*bwms.BaseResponseInventory, error) {
	resp, err := rpc.WmsClient.GetInventory(h.Context, &rpcwms.GetInventoryReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseInventory{Code: codeOK(), Data: toInventory(resp.Inventory), Message: "ok"}, nil
}
func (h *ListInventoryByPageService) Run(req *bwms.InventoryQueryRequest) (*bwms.BaseResponsePageInventory, error) {
	resp, err := rpc.WmsClient.ListInventory(h.Context, &rpcwms.ListInventoryReq{PageNum: req.PageNum, PageSize: req.PageSize, WarehouseId: req.WarehouseId, LocationId: req.LocationId, SkuId: req.SkuId, LotNo: req.LotNo, ExpireDate: req.ExpireDate})
	if err != nil {
		return nil, err
	}
	records := make([]*bwms.Inventory, 0, len(resp.Records))
	for _, rec := range resp.Records {
		records = append(records, toInventory(rec))
	}
	return &bwms.BaseResponsePageInventory{Code: codeOK(), Data: &bwms.PageInventory{Records: records, PageNumber: req.PageNum, PageSize: req.PageSize, TotalPage: toTotalPage(resp.Total, req.PageSize), TotalRow: resp.Total}, Message: "ok"}, nil
}
func (h *CheckInventoryService) Run(req *bwms.CheckInventoryRequest) (*bwms.BaseResponseCheckInventory, error) {
	items := make([]*rpcwms.CheckInventoryItem, 0, len(req.Items))
	for _, it := range req.Items {
		items = append(items, &rpcwms.CheckInventoryItem{WarehouseId: it.WarehouseId, LocationId: it.LocationId, SkuId: it.SkuId, LotNo: it.LotNo, ExpireDate: it.ExpireDate, RequiredQty: it.RequiredQty})
	}
	resp, err := rpc.WmsClient.CheckInventory(h.Context, &rpcwms.CheckInventoryReq{Items: items})
	if err != nil {
		return nil, err
	}
	results := make([]*bwms.CheckInventoryResultItem, 0, len(resp.Results))
	for _, r := range resp.Results {
		results = append(results, &bwms.CheckInventoryResultItem{Requested: &bwms.CheckInventoryItem{WarehouseId: r.Requested.WarehouseId, LocationId: r.Requested.LocationId, SkuId: r.Requested.SkuId, LotNo: r.Requested.LotNo, ExpireDate: r.Requested.ExpireDate, RequiredQty: r.Requested.RequiredQty}, Result: int32(r.Result), AvailableQty: r.AvailableQty})
	}
	return &bwms.BaseResponseCheckInventory{Code: codeOK(), Data: &bwms.CheckInventoryResponse{Results: results, AllPass: resp.AllPass}, Message: "ok"}, nil
}

func (h *CreateInboundOrderService) Run(req *bwms.InboundCreateRequest) (*bwms.BaseResponseLong, error) {
	items := make([]*rpcwms.InboundOrderItemInput, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, &rpcwms.InboundOrderItemInput{SkuId: item.SkuId, LocationId: item.LocationId, LotNo: item.LotNo, ExpireDate: item.ExpireDate, Qty: item.Qty})
	}
	resp, err := rpc.WmsClient.CreateInboundOrder(h.Context, &rpcwms.CreateInboundOrderReq{WarehouseId: req.WarehouseId, SupplierName: req.SupplierName, Remark: req.Remark, OperatorUserId: getUserID(h.Context), Items: items})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseLong{Code: codeOK(), Data: resp.Id, Message: "ok"}, nil
}
func (h *ReceiveInboundOrderService) Run(req *bwms.OrderOperateRequest) (*bwms.BaseResponseBoolean, error) {
	resp, err := rpc.WmsClient.ReceiveInboundOrder(h.Context, &rpcwms.ReceiveInboundOrderReq{Id: req.Id, OperatorUserId: getUserID(h.Context)})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseBoolean{Code: codeOK(), Data: resp.Success, Message: "ok"}, nil
}
func (h *GetInboundOrderByIdService) Run(req *bwms.GetByIdRequest) (*bwms.BaseResponseInboundOrder, error) {
	resp, err := rpc.WmsClient.GetInboundOrder(h.Context, &rpcwms.GetInboundOrderReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseInboundOrder{Code: codeOK(), Data: toInboundOrder(resp.Order), Message: "ok"}, nil
}
func (h *ListInboundOrderByPageService) Run(req *bwms.InboundOrderQueryRequest) (*bwms.BaseResponsePageInboundOrder, error) {
	resp, err := rpc.WmsClient.ListInboundOrder(h.Context, &rpcwms.ListInboundOrderReq{PageNum: req.PageNum, PageSize: req.PageSize, InboundNo: req.InboundNo, WarehouseId: req.WarehouseId, Status: rpcwms.InboundStatus(req.Status)})
	if err != nil {
		return nil, err
	}
	records := make([]*bwms.InboundOrder, 0, len(resp.Records))
	for _, rec := range resp.Records {
		records = append(records, toInboundOrder(rec))
	}
	return &bwms.BaseResponsePageInboundOrder{Code: codeOK(), Data: &bwms.PageInboundOrder{Records: records, PageNumber: req.PageNum, PageSize: req.PageSize, TotalPage: toTotalPage(resp.Total, req.PageSize), TotalRow: resp.Total}, Message: "ok"}, nil
}

func (h *CreateOutboundOrderService) Run(req *bwms.OutboundCreateRequest) (*bwms.BaseResponseLong, error) {
	items := make([]*rpcwms.OutboundOrderItemInput, 0, len(req.Items))
	for _, item := range req.Items {
		items = append(items, &rpcwms.OutboundOrderItemInput{SkuId: item.SkuId, LocationId: item.LocationId, LotNo: item.LotNo, ExpireDate: item.ExpireDate, Qty: item.Qty})
	}
	resp, err := rpc.WmsClient.CreateOutboundOrder(h.Context, &rpcwms.CreateOutboundOrderReq{WarehouseId: req.WarehouseId, CustomerName: req.CustomerName, Remark: req.Remark, OperatorUserId: getUserID(h.Context), Items: items})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseLong{Code: codeOK(), Data: resp.Id, Message: "ok"}, nil
}
func (h *ShipOutboundOrderService) Run(req *bwms.OrderOperateRequest) (*bwms.BaseResponseBoolean, error) {
	resp, err := rpc.WmsClient.ShipOutboundOrder(h.Context, &rpcwms.ShipOutboundOrderReq{Id: req.Id, OperatorUserId: getUserID(h.Context)})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseBoolean{Code: codeOK(), Data: resp.Success, Message: "ok"}, nil
}
func (h *GetOutboundOrderByIdService) Run(req *bwms.GetByIdRequest) (*bwms.BaseResponseOutboundOrder, error) {
	resp, err := rpc.WmsClient.GetOutboundOrder(h.Context, &rpcwms.GetOutboundOrderReq{Id: req.Id})
	if err != nil {
		return nil, err
	}
	return &bwms.BaseResponseOutboundOrder{Code: codeOK(), Data: toOutboundOrder(resp.Order), Message: "ok"}, nil
}
func (h *ListOutboundOrderByPageService) Run(req *bwms.OutboundOrderQueryRequest) (*bwms.BaseResponsePageOutboundOrder, error) {
	resp, err := rpc.WmsClient.ListOutboundOrder(h.Context, &rpcwms.ListOutboundOrderReq{PageNum: req.PageNum, PageSize: req.PageSize, OutboundNo: req.OutboundNo, WarehouseId: req.WarehouseId, Status: rpcwms.OutboundStatus(req.Status)})
	if err != nil {
		return nil, err
	}
	records := make([]*bwms.OutboundOrder, 0, len(resp.Records))
	for _, rec := range resp.Records {
		records = append(records, toOutboundOrder(rec))
	}
	return &bwms.BaseResponsePageOutboundOrder{Code: codeOK(), Data: &bwms.PageOutboundOrder{Records: records, PageNumber: req.PageNum, PageSize: req.PageSize, TotalPage: toTotalPage(resp.Total, req.PageSize), TotalRow: resp.Total}, Message: "ok"}, nil
}

func (h *ListInventoryTxnByPageService) Run(req *bwms.InventoryTxnQueryRequest) (*bwms.BaseResponsePageInventoryTxn, error) {
	resp, err := rpc.WmsClient.ListInventoryTxn(h.Context, &rpcwms.ListInventoryTxnReq{PageNum: req.PageNum, PageSize: req.PageSize, BizType: rpcwms.BizType(req.BizType), BizNo: req.BizNo, WarehouseId: req.WarehouseId, LocationId: req.LocationId, SkuId: req.SkuId})
	if err != nil {
		return nil, err
	}
	records := make([]*bwms.InventoryTxn, 0, len(resp.Records))
	for _, rec := range resp.Records {
		records = append(records, &bwms.InventoryTxn{Id: rec.Id, BizType: int32(rec.BizType), BizNo: rec.BizNo, WarehouseId: rec.WarehouseId, LocationId: rec.LocationId, SkuId: rec.SkuId, LotNo: rec.LotNo, ExpireDate: rec.ExpireDate, QtyChange: rec.QtyChange, BeforeQty: rec.BeforeQty, AfterQty: rec.AfterQty, OperatorUserId: rec.OperatorUserId, CreateTime: rec.CreateTime})
	}
	return &bwms.BaseResponsePageInventoryTxn{Code: codeOK(), Data: &bwms.PageInventoryTxn{Records: records, PageNumber: req.PageNum, PageSize: req.PageSize, TotalPage: toTotalPage(resp.Total, req.PageSize), TotalRow: resp.Total}, Message: "ok"}, nil
}

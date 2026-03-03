package wms

import (
	"context"

	"github.com/MoScenix/wms/app/bff/biz/service"
	"github.com/MoScenix/wms/app/bff/biz/utils"
	wms "github.com/MoScenix/wms/app/bff/hertz_gen/bff/wms"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// AddWarehouse .
// @router /wms/warehouse/add [POST]
func AddWarehouse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.WarehouseAddRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseLong{}
	resp, err = service.NewAddWarehouseService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateWarehouse .
// @router /wms/warehouse/update [POST]
func UpdateWarehouse(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.WarehouseUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseBoolean{}
	resp, err = service.NewUpdateWarehouseService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetWarehouseById .
// @router /wms/warehouse/get [GET]
func GetWarehouseById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.GetByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseWarehouse{}
	resp, err = service.NewGetWarehouseByIdService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListWarehouseByPage .
// @router /wms/warehouse/list/page [POST]
func ListWarehouseByPage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.WarehouseQueryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponsePageWarehouse{}
	resp, err = service.NewListWarehouseByPageService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddLocation .
// @router /wms/location/add [POST]
func AddLocation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.LocationAddRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseLong{}
	resp, err = service.NewAddLocationService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateLocation .
// @router /wms/location/update [POST]
func UpdateLocation(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.LocationUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseBoolean{}
	resp, err = service.NewUpdateLocationService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetLocationById .
// @router /wms/location/get [GET]
func GetLocationById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.GetByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseLocation{}
	resp, err = service.NewGetLocationByIdService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListLocationByPage .
// @router /wms/location/list/page [POST]
func ListLocationByPage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.LocationQueryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponsePageLocation{}
	resp, err = service.NewListLocationByPageService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// AddSku .
// @router /wms/sku/add [POST]
func AddSku(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.SkuAddRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseLong{}
	resp, err = service.NewAddSkuService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// UpdateSku .
// @router /wms/sku/update [POST]
func UpdateSku(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.SkuUpdateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseBoolean{}
	resp, err = service.NewUpdateSkuService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetSkuById .
// @router /wms/sku/get [GET]
func GetSkuById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.GetByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseSku{}
	resp, err = service.NewGetSkuByIdService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListSkuByPage .
// @router /wms/sku/list/page [POST]
func ListSkuByPage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.SkuQueryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponsePageSku{}
	resp, err = service.NewListSkuByPageService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetInventoryById .
// @router /wms/inventory/get [GET]
func GetInventoryById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.GetByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseInventory{}
	resp, err = service.NewGetInventoryByIdService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListInventoryByPage .
// @router /wms/inventory/list/page [POST]
func ListInventoryByPage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.InventoryQueryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponsePageInventory{}
	resp, err = service.NewListInventoryByPageService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CheckInventory .
// @router /wms/inventory/check [POST]
func CheckInventory(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.CheckInventoryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseCheckInventory{}
	resp, err = service.NewCheckInventoryService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreateInboundOrder .
// @router /wms/inbound/create [POST]
func CreateInboundOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.InboundCreateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseLong{}
	resp, err = service.NewCreateInboundOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ReceiveInboundOrder .
// @router /wms/inbound/receive [POST]
func ReceiveInboundOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.OrderOperateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseBoolean{}
	resp, err = service.NewReceiveInboundOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetInboundOrderById .
// @router /wms/inbound/get [GET]
func GetInboundOrderById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.GetByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseInboundOrder{}
	resp, err = service.NewGetInboundOrderByIdService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListInboundOrderByPage .
// @router /wms/inbound/list/page [POST]
func ListInboundOrderByPage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.InboundOrderQueryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponsePageInboundOrder{}
	resp, err = service.NewListInboundOrderByPageService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// CreateOutboundOrder .
// @router /wms/outbound/create [POST]
func CreateOutboundOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.OutboundCreateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseLong{}
	resp, err = service.NewCreateOutboundOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ShipOutboundOrder .
// @router /wms/outbound/ship [POST]
func ShipOutboundOrder(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.OrderOperateRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseBoolean{}
	resp, err = service.NewShipOutboundOrderService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// GetOutboundOrderById .
// @router /wms/outbound/get [GET]
func GetOutboundOrderById(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.GetByIdRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponseOutboundOrder{}
	resp, err = service.NewGetOutboundOrderByIdService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListOutboundOrderByPage .
// @router /wms/outbound/list/page [POST]
func ListOutboundOrderByPage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.OutboundOrderQueryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponsePageOutboundOrder{}
	resp, err = service.NewListOutboundOrderByPageService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// ListInventoryTxnByPage .
// @router /wms/txn/list/page [POST]
func ListInventoryTxnByPage(ctx context.Context, c *app.RequestContext) {
	var err error
	var req wms.InventoryTxnQueryRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp := &wms.BaseResponsePageInventoryTxn{}
	resp, err = service.NewListInventoryTxnByPageService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

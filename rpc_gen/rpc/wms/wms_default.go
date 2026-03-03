package wms

import (
	"context"
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func AddWarehouse(ctx context.Context, req *wms.AddWarehouseReq, callOptions ...callopt.Option) (resp *wms.AddWarehouseResp, err error) {
	resp, err = defaultClient.AddWarehouse(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddWarehouse call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateWarehouse(ctx context.Context, req *wms.UpdateWarehouseReq, callOptions ...callopt.Option) (resp *wms.UpdateWarehouseResp, err error) {
	resp, err = defaultClient.UpdateWarehouse(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateWarehouse call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetWarehouse(ctx context.Context, req *wms.GetWarehouseReq, callOptions ...callopt.Option) (resp *wms.GetWarehouseResp, err error) {
	resp, err = defaultClient.GetWarehouse(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetWarehouse call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListWarehouse(ctx context.Context, req *wms.ListWarehouseReq, callOptions ...callopt.Option) (resp *wms.ListWarehouseResp, err error) {
	resp, err = defaultClient.ListWarehouse(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListWarehouse call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddLocation(ctx context.Context, req *wms.AddLocationReq, callOptions ...callopt.Option) (resp *wms.AddLocationResp, err error) {
	resp, err = defaultClient.AddLocation(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddLocation call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateLocation(ctx context.Context, req *wms.UpdateLocationReq, callOptions ...callopt.Option) (resp *wms.UpdateLocationResp, err error) {
	resp, err = defaultClient.UpdateLocation(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateLocation call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetLocation(ctx context.Context, req *wms.GetLocationReq, callOptions ...callopt.Option) (resp *wms.GetLocationResp, err error) {
	resp, err = defaultClient.GetLocation(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetLocation call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListLocation(ctx context.Context, req *wms.ListLocationReq, callOptions ...callopt.Option) (resp *wms.ListLocationResp, err error) {
	resp, err = defaultClient.ListLocation(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListLocation call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func AddSku(ctx context.Context, req *wms.AddSkuReq, callOptions ...callopt.Option) (resp *wms.AddSkuResp, err error) {
	resp, err = defaultClient.AddSku(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "AddSku call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func UpdateSku(ctx context.Context, req *wms.UpdateSkuReq, callOptions ...callopt.Option) (resp *wms.UpdateSkuResp, err error) {
	resp, err = defaultClient.UpdateSku(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "UpdateSku call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetSku(ctx context.Context, req *wms.GetSkuReq, callOptions ...callopt.Option) (resp *wms.GetSkuResp, err error) {
	resp, err = defaultClient.GetSku(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetSku call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListSku(ctx context.Context, req *wms.ListSkuReq, callOptions ...callopt.Option) (resp *wms.ListSkuResp, err error) {
	resp, err = defaultClient.ListSku(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListSku call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetInventory(ctx context.Context, req *wms.GetInventoryReq, callOptions ...callopt.Option) (resp *wms.GetInventoryResp, err error) {
	resp, err = defaultClient.GetInventory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetInventory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListInventory(ctx context.Context, req *wms.ListInventoryReq, callOptions ...callopt.Option) (resp *wms.ListInventoryResp, err error) {
	resp, err = defaultClient.ListInventory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListInventory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CheckInventory(ctx context.Context, req *wms.CheckInventoryReq, callOptions ...callopt.Option) (resp *wms.CheckInventoryResp, err error) {
	resp, err = defaultClient.CheckInventory(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CheckInventory call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateInboundOrder(ctx context.Context, req *wms.CreateInboundOrderReq, callOptions ...callopt.Option) (resp *wms.CreateInboundOrderResp, err error) {
	resp, err = defaultClient.CreateInboundOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateInboundOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ReceiveInboundOrder(ctx context.Context, req *wms.ReceiveInboundOrderReq, callOptions ...callopt.Option) (resp *wms.ReceiveInboundOrderResp, err error) {
	resp, err = defaultClient.ReceiveInboundOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ReceiveInboundOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetInboundOrder(ctx context.Context, req *wms.GetInboundOrderReq, callOptions ...callopt.Option) (resp *wms.GetInboundOrderResp, err error) {
	resp, err = defaultClient.GetInboundOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetInboundOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListInboundOrder(ctx context.Context, req *wms.ListInboundOrderReq, callOptions ...callopt.Option) (resp *wms.ListInboundOrderResp, err error) {
	resp, err = defaultClient.ListInboundOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListInboundOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func CreateOutboundOrder(ctx context.Context, req *wms.CreateOutboundOrderReq, callOptions ...callopt.Option) (resp *wms.CreateOutboundOrderResp, err error) {
	resp, err = defaultClient.CreateOutboundOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "CreateOutboundOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ShipOutboundOrder(ctx context.Context, req *wms.ShipOutboundOrderReq, callOptions ...callopt.Option) (resp *wms.ShipOutboundOrderResp, err error) {
	resp, err = defaultClient.ShipOutboundOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ShipOutboundOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetOutboundOrder(ctx context.Context, req *wms.GetOutboundOrderReq, callOptions ...callopt.Option) (resp *wms.GetOutboundOrderResp, err error) {
	resp, err = defaultClient.GetOutboundOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetOutboundOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListOutboundOrder(ctx context.Context, req *wms.ListOutboundOrderReq, callOptions ...callopt.Option) (resp *wms.ListOutboundOrderResp, err error) {
	resp, err = defaultClient.ListOutboundOrder(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListOutboundOrder call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func ListInventoryTxn(ctx context.Context, req *wms.ListInventoryTxnReq, callOptions ...callopt.Option) (resp *wms.ListInventoryTxnResp, err error) {
	resp, err = defaultClient.ListInventoryTxn(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "ListInventoryTxn call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

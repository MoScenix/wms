import request from '@/request'

export const addWarehouse = (data: any) => request.post('/wms/warehouse/add', data)
export const updateWarehouse = (data: any) => request.post('/wms/warehouse/update', data)
export const getWarehouseById = (params: any) => request.get('/wms/warehouse/get', { params })
export const listWarehouseByPage = (data: any) => request.post('/wms/warehouse/list/page', data)

export const addLocation = (data: any) => request.post('/wms/location/add', data)
export const updateLocation = (data: any) => request.post('/wms/location/update', data)
export const getLocationById = (params: any) => request.get('/wms/location/get', { params })
export const listLocationByPage = (data: any) => request.post('/wms/location/list/page', data)

export const addSku = (data: any) => request.post('/wms/sku/add', data)
export const updateSku = (data: any) => request.post('/wms/sku/update', data)
export const getSkuById = (params: any) => request.get('/wms/sku/get', { params })
export const listSkuByPage = (data: any) => request.post('/wms/sku/list/page', data)

export const getInventoryById = (params: any) => request.get('/wms/inventory/get', { params })
export const listInventoryByPage = (data: any) => request.post('/wms/inventory/list/page', data)
export const checkInventory = (data: any) => request.post('/wms/inventory/check', data)

export const createInboundOrder = (data: any) => request.post('/wms/inbound/create', data)
export const receiveInboundOrder = (data: any) => request.post('/wms/inbound/receive', data)
export const getInboundOrderById = (params: any) => request.get('/wms/inbound/get', { params })
export const listInboundOrderByPage = (data: any) => request.post('/wms/inbound/list/page', data)

export const createOutboundOrder = (data: any) => request.post('/wms/outbound/create', data)
export const shipOutboundOrder = (data: any) => request.post('/wms/outbound/ship', data)
export const getOutboundOrderById = (params: any) => request.get('/wms/outbound/get', { params })
export const listOutboundOrderByPage = (data: any) => request.post('/wms/outbound/list/page', data)

export const listInventoryTxnByPage = (data: any) => request.post('/wms/txn/list/page', data)

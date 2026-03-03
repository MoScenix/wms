package model

import (
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"gorm.io/gorm"
)

type Location struct {
	gorm.Model
	WarehouseID  uint   `gorm:"index;not null"`
	LocationCode string `gorm:"type:varchar(64);not null;index:idx_wh_loc_code,unique"`
	LocationName string `gorm:"type:varchar(128);not null"`
	LocationType string `gorm:"type:varchar(32);default:''"`
}

func (Location) TableName() string { return "location" }

func toLocationInfo(m Location) *wms.LocationInfo {
	return &wms.LocationInfo{Id: int64(m.ID), WarehouseId: int64(m.WarehouseID), LocationCode: m.LocationCode, LocationName: m.LocationName, LocationType: m.LocationType, CreateTime: toTimeString(m.CreatedAt), UpdateTime: toTimeString(m.UpdatedAt)}
}

func (p *WmsProQuery) AddLocation(req *wms.AddLocationReq) (*wms.AddLocationResp, error) {
	m := Location{WarehouseID: uint(req.WarehouseId), LocationCode: req.LocationCode, LocationName: req.LocationName, LocationType: req.LocationType}
	if err := p.q.db.WithContext(p.q.ctx).Create(&m).Error; err != nil {
		return nil, err
	}
	return &wms.AddLocationResp{Id: int64(m.ID)}, nil
}

func (p *WmsProQuery) UpdateLocation(req *wms.UpdateLocationReq) (*wms.UpdateLocationResp, error) {
	updates := map[string]any{}
	if req.LocationName != "" {
		updates["location_name"] = req.LocationName
	}
	if req.LocationType != "" {
		updates["location_type"] = req.LocationType
	}
	if len(updates) == 0 {
		return &wms.UpdateLocationResp{Success: true}, nil
	}
	if err := p.q.db.WithContext(p.q.ctx).Model(&Location{}).Where("id = ?", req.Id).Updates(updates).Error; err != nil {
		return nil, err
	}
	p.cacheDel(p.key("location", req.Id))
	return &wms.UpdateLocationResp{Success: true}, nil
}

func (p *WmsProQuery) GetLocation(id int64) (*wms.GetLocationResp, error) {
	cacheKey := p.key("location", id)
	if cached := new(wms.LocationInfo); p.cacheGet(cacheKey, cached) {
		return &wms.GetLocationResp{Location: cached}, nil
	}
	var m Location
	if err := p.q.db.WithContext(p.q.ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	data := toLocationInfo(m)
	p.cacheSet(cacheKey, data)
	return &wms.GetLocationResp{Location: data}, nil
}

func (p *WmsProQuery) ListLocation(req *wms.ListLocationReq) (*wms.ListLocationResp, error) {
	pageNum, pageSize := pageNormalize(req.PageNum, req.PageSize)
	q := p.q.db.WithContext(p.q.ctx).Model(&Location{}).
		Where("location_code like ?", likePrefix(req.LocationCode)).
		Where("location_name like ?", likePrefix(req.LocationName))
	if req.WarehouseId > 0 {
		q = q.Where("warehouse_id = ?", req.WarehouseId)
	}
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}
	var rows []Location
	if err := q.Order("id desc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}
	ret := make([]*wms.LocationInfo, 0, len(rows))
	for _, it := range rows {
		ret = append(ret, toLocationInfo(it))
	}
	return &wms.ListLocationResp{Records: ret, Total: total}, nil
}

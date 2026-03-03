package model

import (
	wms "github.com/MoScenix/wms/rpc_gen/kitex_gen/wms"
	"gorm.io/gorm"
)

type Sku struct {
	gorm.Model
	SkuCode  string `gorm:"type:varchar(64);not null;uniqueIndex"`
	SkuName  string `gorm:"type:varchar(128);not null"`
	Unit     string `gorm:"type:varchar(32);default:''"`
	Category string `gorm:"type:varchar(64);default:''"`
}

func (Sku) TableName() string { return "sku" }

func toSkuInfo(m Sku) *wms.SkuInfo {
	return &wms.SkuInfo{Id: int64(m.ID), SkuCode: m.SkuCode, SkuName: m.SkuName, Unit: m.Unit, Category: m.Category, CreateTime: toTimeString(m.CreatedAt), UpdateTime: toTimeString(m.UpdatedAt)}
}

func (p *WmsProQuery) AddSku(req *wms.AddSkuReq) (*wms.AddSkuResp, error) {
	m := Sku{SkuCode: req.SkuCode, SkuName: req.SkuName, Unit: req.Unit, Category: req.Category}
	if err := p.q.db.WithContext(p.q.ctx).Create(&m).Error; err != nil {
		return nil, err
	}
	return &wms.AddSkuResp{Id: int64(m.ID)}, nil
}

func (p *WmsProQuery) UpdateSku(req *wms.UpdateSkuReq) (*wms.UpdateSkuResp, error) {
	updates := map[string]any{}
	if req.SkuName != "" {
		updates["sku_name"] = req.SkuName
	}
	if req.Unit != "" {
		updates["unit"] = req.Unit
	}
	if req.Category != "" {
		updates["category"] = req.Category
	}
	if len(updates) == 0 {
		return &wms.UpdateSkuResp{Success: true}, nil
	}
	if err := p.q.db.WithContext(p.q.ctx).Model(&Sku{}).Where("id = ?", req.Id).Updates(updates).Error; err != nil {
		return nil, err
	}
	p.cacheDel(p.key("sku", req.Id))
	return &wms.UpdateSkuResp{Success: true}, nil
}

func (p *WmsProQuery) GetSku(id int64) (*wms.GetSkuResp, error) {
	cacheKey := p.key("sku", id)
	if cached := new(wms.SkuInfo); p.cacheGet(cacheKey, cached) {
		return &wms.GetSkuResp{Sku: cached}, nil
	}
	var m Sku
	if err := p.q.db.WithContext(p.q.ctx).First(&m, id).Error; err != nil {
		return nil, err
	}
	data := toSkuInfo(m)
	p.cacheSet(cacheKey, data)
	return &wms.GetSkuResp{Sku: data}, nil
}

func (p *WmsProQuery) ListSku(req *wms.ListSkuReq) (*wms.ListSkuResp, error) {
	pageNum, pageSize := pageNormalize(req.PageNum, req.PageSize)
	q := p.q.db.WithContext(p.q.ctx).Model(&Sku{}).
		Where("sku_code like ?", likePrefix(req.SkuCode)).
		Where("sku_name like ?", likePrefix(req.SkuName))
	var total int64
	if err := q.Count(&total).Error; err != nil {
		return nil, err
	}
	var rows []Sku
	if err := q.Order("id desc").Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&rows).Error; err != nil {
		return nil, err
	}
	ret := make([]*wms.SkuInfo, 0, len(rows))
	for _, it := range rows {
		ret = append(ret, toSkuInfo(it))
	}
	return &wms.ListSkuResp{Records: ret, Total: total}, nil
}

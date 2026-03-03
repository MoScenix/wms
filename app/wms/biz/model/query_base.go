package model

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type WmsQuery struct {
	ctx context.Context
	db  *gorm.DB
}

func NewWmsQuery(ctx context.Context, db *gorm.DB) *WmsQuery { return &WmsQuery{ctx: ctx, db: db} }

type WmsProQuery struct {
	q      *WmsQuery
	rdb    *redis.Client
	prefix string
}

func NewWmsProQuery(ctx context.Context, db *gorm.DB, rdb *redis.Client) *WmsProQuery {
	return &WmsProQuery{q: NewWmsQuery(ctx, db), rdb: rdb, prefix: "wms"}
}

func toTimeString(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.Format(time.RFC3339)
}

func pageNormalize(pageNum, pageSize int64) (int64, int64) {
	if pageNum <= 0 {
		pageNum = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return pageNum, pageSize
}

func likePrefix(s string) string {
	if s == "" {
		return "%"
	}
	return strings.TrimSpace(s) + "%"
}

func (p *WmsProQuery) key(kind string, id int64) string {
	return fmt.Sprintf("%s_%s_%d", p.prefix, kind, id)
}

func (p *WmsProQuery) cacheGet(key string, v any) bool {
	if p.rdb == nil {
		return false
	}
	val, err := p.rdb.Get(p.q.ctx, key).Result()
	if err != nil || val == "" {
		return false
	}
	return json.Unmarshal([]byte(val), v) == nil
}

func (p *WmsProQuery) cacheSet(key string, v any) {
	if p.rdb == nil {
		return
	}
	if b, err := json.Marshal(v); err == nil {
		_ = p.rdb.Set(p.q.ctx, key, b, time.Hour).Err()
	}
}

func (p *WmsProQuery) cacheDel(key string) {
	if p.rdb != nil {
		_ = p.rdb.Del(p.q.ctx, key).Err()
	}
}

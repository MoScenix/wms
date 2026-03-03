package middleware

import (
	"context"
	"fmt"
	"strings"

	"github.com/MoScenix/wms/app/bff/biz/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		session := sessions.Default(c)
		fmt.Println(c.Path())
		userID := session.Get(utils.UserIdKey)
		userRole := session.Get(utils.UserRoleKey)
		if userID == nil || userRole == nil {
			c.Next(ctx)
			return
		}
		ctx = context.WithValue(ctx, utils.UserIdKey, userID)
		ctx = context.WithValue(ctx, utils.UserRoleKey, userRole)
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		path := string(c.Path())
		userID := ctx.Value(utils.UserIdKey)
		userRole := ctx.Value(utils.UserRoleKey)

		if strings.Contains(path, "/wms/") {
			if userID == nil {
				c.String(consts.StatusUnauthorized, "please login first")
				c.Abort()
				return
			}
			if strings.Contains(path, "/warehouse/") || strings.Contains(path, "/location/") || strings.Contains(path, "/sku/") {
				if userRole != utils.AdminRole {
					c.String(consts.StatusForbidden, "admin role required")
					c.Abort()
					return
				}
			}
		}
		c.Next(ctx)
	}
}

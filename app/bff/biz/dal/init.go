package dal

import (
	"github.com/MoScenix/wms/app/bff/biz/dal/mysql"
	"github.com/MoScenix/wms/app/bff/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

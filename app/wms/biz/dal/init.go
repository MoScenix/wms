package dal

import (
	"github.com/MoScenix/wms/app/wms/biz/dal/mysql"
	"github.com/MoScenix/wms/app/wms/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

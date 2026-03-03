package dal

import (
	"github.com/MoScenix/wms/app/user/biz/dal/mysql"
	"github.com/MoScenix/wms/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}

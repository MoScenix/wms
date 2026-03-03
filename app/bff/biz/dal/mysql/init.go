package mysql

import (
	"fmt"
	"os"
	"strings"

	"github.com/MoScenix/wms/app/bff/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Init() {
	rawDSN := conf.GetConf().MySQL.DSN
	dsn := rawDSN
	if strings.Contains(rawDSN, "%s") {
		dsn = fmt.Sprintf(rawDSN, os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	}

	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
}

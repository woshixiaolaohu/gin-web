package initialize

import (
	"gin-vue-admin/global"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DBType {
	case "mysql":
		return GormMysql()
	case "oracle":
		return GormOracle()
	case "mssql":
		return GormMssql()
	case "pgsql":
		return GormPgsql()
	case "sqlite":
		return GoSqlite()
	default:
		return GormMysql()
	}
}

package initialize

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"os"
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

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		system.SysApi{},
		system.SysAuthority{},
		system.SysUser{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}

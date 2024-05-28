package initialize

import (
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gorm.io/gorm"
)

const sys = "system"

func DBList() {
	dbMap := make(map[string]*gorm.DB)
	for _, info := range global.GVA_CONFIG.DBList {
		if info.Disable {
			continue
		}
		switch info.Type {
		case "mysql":
			dbMap[info.AliasName] = GormMysqlByConfig(config.Mysql{GeneralDB: info.GeneralDB})
		case "mssql":
			dbMap[info.AliasName] = GormMssqlByConfig(config.Mssql{GeneralDB: info.GeneralDB})
		case "pgsql":
			dbMap[info.AliasName] = GormPgsqlByConfig(config.Pgsql{GeneralDB: info.GeneralDB})
		case "oracle":
			dbMap[info.AliasName] = GormOracleByConfig(config.Oracle{GeneralDB: info.GeneralDB})
		case "sqlite":
			dbMap[info.AliasName] = GormSqliteByConfig(config.Sqlite{GeneralDB: info.GeneralDB})
		default:
			continue
		}
	}
	if sysDB, ok := dbMap[sys]; ok {
		global.GVA_DB = sysDB
	}
	global.GVA_DBList = dbMap
}

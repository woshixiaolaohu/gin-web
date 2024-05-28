package initialize

import (
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize/internal"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

// GormMssql 初始化 Mssql 数据库
func GormMssql() *gorm.DB {
	m := global.GVA_CONFIG.Mssql
	if m.DBName == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(),
		DefaultStringSize: 191,
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

// GormMssqlByConfig 初始化 Mssql 数据库，根据传出的配置
func GormMssqlByConfig(m config.Mssql) *gorm.DB {
	if m.DBName == "" {
		return nil
	}
	mssqlConfig := sqlserver.Config{
		DSN:               m.Dsn(),
		DefaultStringSize: 191,
	}
	if db, err := gorm.Open(sqlserver.New(mssqlConfig), internal.Gorm.Config(m.Prefix, m.Singular)); err != nil {
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+m.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

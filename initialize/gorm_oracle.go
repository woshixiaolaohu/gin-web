package initialize

import (
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize/internal"
	"github.com/dzwvip/oracle"
	"gorm.io/gorm"
)

// GormOracle 初始化Oracle数据库
func GormOracle() *gorm.DB {
	o := global.GVA_CONFIG.Oracle
	if o.DBName == "" {
		return nil
	}
	oracleConfig := oracle.Config{
		DSN:               o.Dsn(),
		DefaultStringSize: 191,
	}
	if db, err := gorm.Open(oracle.New(oracleConfig), internal.Gorm.Config(o.Prefix, o.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(o.MaxIdleConns)
		sqlDB.SetMaxOpenConns(o.MaxOpenConns)
		return db
	}
}

// GormOracleByConfig 初始化Oracle数据库，使用传入配置
func GormOracleByConfig(o config.Oracle) *gorm.DB {
	if o.DBName == "" {
		return nil
	}
	oracleConfig := oracle.Config{
		DSN:               o.Dsn(),
		DefaultStringSize: 191,
	}
	if db, err := gorm.Open(oracle.New(oracleConfig), internal.Gorm.Config(o.Prefix, o.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(o.MaxIdleConns)
		sqlDB.SetMaxOpenConns(o.MaxOpenConns)
		return db
	}
}

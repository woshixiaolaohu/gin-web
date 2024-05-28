package initialize

import (
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize/internal"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// GoSqlite 初始化Sqlite数据库
func GoSqlite() *gorm.DB {
	s := global.GVA_CONFIG.Sqlite
	if s.DBName == "" {
		return nil
	}
	if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config(s.Prefix, s.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}

// GormSqliteByConfig 初始化 Sqlite 数据库，根据传入配置
func GormSqliteByConfig(s config.Sqlite) *gorm.DB {
	if s.DBName == "" {
		return nil
	}
	if db, err := gorm.Open(sqlite.Open(s.Dsn()), internal.Gorm.Config(s.Prefix, s.Singular)); err != nil {
		panic(err)
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(s.MaxIdleConns)
		sqlDB.SetMaxOpenConns(s.MaxOpenConns)
		return db
	}
}

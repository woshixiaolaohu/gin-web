package initialize

import (
	"gin-vue-admin/config"
	"gorm.io/gorm"
)

// GormPgsql 初始化 Pgsql 数据库
func GormPgsql() *gorm.DB {
	return nil
	//p := global.GVA_CONFIG.Pgsql
	//if p.DBName == "" {
	//	return nil
	//}
	//pgsqlConfig := postgres.Config{
	//	DSN:                  p.Dsn(),
	//	PreferSimpleProtocol: false,
	//}
	//if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular)); err != nil {
	//	return nil
	//} else {
	//	sqlDB, _ := db.DB()
	//	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	//	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	//	return db
	//}
}

// GormPgsqlByConfig 初始化 Pgsql 数据库，根据传入配置
func GormPgsqlByConfig(p config.Pgsql) *gorm.DB {
	return nil
	//if p.DBName == "" {
	//	return nil
	//}
	//pgsqlConfig := postgres.Config{
	//	DSN:                  p.Dsn(),
	//	PreferSimpleProtocol: false,
	//}
	//if db, err := gorm.Open(postgres.New(pgsqlConfig), internal.Gorm.Config(p.Prefix, p.Singular)); err != nil {
	//	return nil
	//} else {
	//	sqlDB, _ := db.DB()
	//	sqlDB.SetMaxIdleConns(p.MaxIdleConns)
	//	sqlDB.SetMaxOpenConns(p.MaxOpenConns)
	//	return db
	//}
}

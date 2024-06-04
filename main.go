package main

import (
	"fmt"
	"gin-vue-admin/cmd"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/zap"
)

// @title 接口文档
// @version 1.0
// @description ginProject
func main() {
	fmt.Println("Hello,GO!")
	cmd.Execute()
	// 初始化Viper 读取配置文件
	global.GVA_VP = cmd.Viper()
	// 初始化其他
	initialize.OtherInit()
	// 初始化zap日志库
	global.GVA_LOG = cmd.Zap()
	// 将logger设置为全局logger
	zap.ReplaceGlobals(global.GVA_LOG)
	// 初始化gorm连接数据库
	global.GVA_DB = initialize.Gorm()
	// 初始化定时任务
	initialize.Timer()
	// 根据配置初始化数据库
	initialize.DBList()
	if global.GVA_DB != nil {
		// 初始化表
		initialize.RegisterTables()
		// 程序结束前关闭数据库连接
		db, _ := global.GVA_DB.DB()
		defer db.Close()
	}

}

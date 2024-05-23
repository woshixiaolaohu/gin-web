package main

import (
	"fmt"
	"gin-vue-admin/cmd"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
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
	zap.ReplaceGlobals(global.GVA_LOG)
	// 初始化gorm连接数据库

}

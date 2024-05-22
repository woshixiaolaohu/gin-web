package main

import (
	"fmt"
	"gin-vue-admin/cmd"
	"gin-vue-admin/global"
)

// @title 接口文档
// @version 1.0
// @description ginProject
func main() {
	fmt.Println("Hello,GO!")
	cmd.Execute()
	// 初始化Viper 读取配置文件
	global.GVA_VP = cmd.Viper()
	//
}

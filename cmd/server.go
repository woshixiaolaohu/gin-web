package cmd

import (
	"fmt"
	"gin-vue-admin/global"
)

type server interface {
	ListenAndServe() error
}

func Execute() {
	//s := initServer("8888")
	fmt.Println("Execute")
}

func RunWindowsServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务

	}
}

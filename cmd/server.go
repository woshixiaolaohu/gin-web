package cmd

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/initialize"
	"gin-vue-admin/service/system"
	"go.uber.org/zap"
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
		initialize.Redis()
	}
	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从 db 加载 jwt 数据
	if global.GVA_DB != nil {
		system.LoadAll()
	}
	
}

package cmd

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	"go.uber.org/zap"
	"os"
)

// Zap 获取zap.Logger
func Zap(logger *zap.Logger) {
	// 判断是否有Director文件夹
	if ok, _ := utils.PathExists(global.GVA_CONFIG.Zap.Director); !ok {
		fmt.Printf("create %v directory\n", global.GVA_CONFIG.Zap.Director)
		_ = os.Mkdir(global.GVA_CONFIG.Zap.Director, os.ModePerm)
	}
}

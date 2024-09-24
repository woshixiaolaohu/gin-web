package initialize

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/plugin/email"
	"gin-vue-admin/utils/plugin"
	"github.com/gin-gonic/gin"
)

func PluginInit(group *gin.RouterGroup, Plugin ...plugin.Plugin) {
	for i := range Plugin {
		PluginGroup := group.Group(Plugin[i].RouterPath())
		Plugin[i].Register(PluginGroup)
	}
}

func InstallPlugin(PrivateGroup *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	fmt.Println("无鉴权插件安装-->", PublicRouter)
	fmt.Println("鉴权插件安装", PrivateGroup)

	PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	// 添加跟角色挂钩权限的插件
	PluginInit(PrivateGroup, email.CreateEmailPlug(
		global.GVA_CONFIG.Email.To,
		global.GVA_CONFIG.Email.From,
		global.GVA_CONFIG.Email.Host,
		global.GVA_CONFIG.Email.Secret,
		global.GVA_CONFIG.Email.NickName,
		global.GVA_CONFIG.Email.Port,
		global.GVA_CONFIG.Email.IsSSL,
	))
}

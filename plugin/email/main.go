package email

import (
	"gin-vue-admin/plugin/email/global"
	"gin-vue-admin/plugin/email/router"
	"github.com/gin-gonic/gin"
)

type emailPlugin struct{}

func CreateEmailPlug(To, From, Host, Secret, NickName string, Port int, IsSSL bool) *emailPlugin {
	global.GlobalConfig.To = To
	global.GlobalConfig.From = From
	global.GlobalConfig.Host = Host
	global.GlobalConfig.Secret = Secret
	global.GlobalConfig.NickName = NickName
	global.GlobalConfig.Port = Port
	global.GlobalConfig.IsSSL = IsSSL
	return &emailPlugin{}
}

func (s *emailPlugin) Register(group *gin.RouterGroup) {
	router.RouterGroupApp.InitEmailRouter(group)
}

func (s *emailPlugin) RouterPath() string {
	return "email"
}

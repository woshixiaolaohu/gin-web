package system

import (
	v1 "gin-vue-admin/api"
	"github.com/gin-gonic/gin"
)

type InitRouter struct{}

func (s *InitRouter) InitInitRouter(Router *gin.RouterGroup) {
	initRouter := Router.Group("init")
	dbApi := v1.ApiGroupApp.SystemApiGroup.DBApi
	{
		initRouter.POST("initDB", dbApi.InitDB)   // 初始化数据库
		initRouter.POST("checkDB", dbApi.CheckDB) // 检查是否需要初始化数据库
	}
}

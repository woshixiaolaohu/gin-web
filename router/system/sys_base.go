package system

import "github.com/gin-gonic/gin"
import v1 "gin-vue-admin/api"

type BaseRouter struct{}

func (b *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRouter) {
	baseRouter := Router.Group("base")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		baseRouter.POST("login", baseApi.Login)
	}
}

package system

import (
	v1 "gin-vue-admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct{}

func (s *ApiRouter) InitApiRouter(Router *gin.RouterGroup, RouterPub *gin.RouterGroup) {
	apiRouter := Router.Group("api").Use(middleware.OperationRecord())
	apiRouterWithoutRecord := Router.Group("api")
	apiPublicRouterWithoutRecord := RouterPub.Group("api")
	apiRouterApi := v1.ApiGroupApp.SystemApiGroup.SystemApiApi
	{
		apiRouter.POST("createApi", apiRouterApi.CreateApi)             // 创建Api
		apiRouter.POST("deleteApi", apiRouterApi.DeleteApi)             // 删除Api
		apiRouter.POST("getApiById", apiRouterApi.GetApiByID)           // 获取单条Api消息
		apiRouter.POST("updateApi", apiRouterApi.UpdateApi)             // 更新Api
		apiRouter.POST("deleteApisByIds", apiRouterApi.DeleteApisByIds) // 删除选中Api
	}
	{
		apiRouterWithoutRecord.POST("getAllApis", apiRouterApi.GetAllApis) // 获取所有Api
		apiRouterWithoutRecord.POST("getApiList", apiRouterApi.GetApiList) // 获取Api列表

	}
	{
		apiPublicRouterWithoutRecord.POST("freshCasbin", apiRouterApi.FreshCasbin) // 刷新Casbin权限
	}
}

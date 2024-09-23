package system

import (
	v1 "gin-vue-admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type DictionaryRouter struct{}

func (c *DictionaryRouter) InitSysDictionaryRouter(Router *gin.RouterGroup) {
	sysDictionaryRouter := Router.Group("sysDictionary").Use(middleware.OperationRecord())
	sysDictionaryRouterWithoutRecord := Router.Group("sysDictionary")
	sysDictionaryApi := v1.ApiGroupApp.SystemApiGroup.DictionaryApi
	{
		sysDictionaryRouter.POST("createSysDictionary", sysDictionaryApi.CreateSysDictionary)   // 新建系统字典
		sysDictionaryRouter.DELETE("deleteSysDictionary", sysDictionaryApi.DeleteSysDictionary) // 删除系统字典
		sysDictionaryRouter.PUT("updateSysDictionary", sysDictionaryApi.UpdateSysDictionary)    // 更新系统字典
	}
	{
		sysDictionaryRouterWithoutRecord.GET("getSysDictionary", sysDictionaryApi.GetSysDictionary)         // 根据 ID 获取系统字典
		sysDictionaryRouterWithoutRecord.GET("getSysDictionaryList", sysDictionaryApi.GetSysDictionaryList) // 获取字典列表
	}
}

package system

import (
	v1 "gin-vue-admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type OperationRecordRouter struct{}

func (s *OperationRecordRouter) InitSysOperationRecordRouter(Router *gin.RouterGroup) {
	operationRecordRouter := Router.Group("sysOperationRecord").Use(middleware.OperationRecord())
	operationRecordRouterWithoutRecord := Router.Group("sysOperationRecord")
	operationRecordApi := v1.ApiGroupApp.SystemApiGroup.OperationRecordApi
	{
		operationRecordRouter.POST("createSysOperationRecord", operationRecordApi.CreateSysOperationRecord)             // 新建 SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecord", operationRecordApi.DeleteSysOperationRecord)           // 删除 SysOperationRecord
		operationRecordRouter.DELETE("deleteSysOperationRecordByIds", operationRecordApi.DeleteSysOperationRecordByIds) // 批量删除 SysOperationRecord
	}
	{
		operationRecordRouterWithoutRecord.GET("getSysOperationRecord", operationRecordApi.GetSysOperationRecord)         // 根据 ID 查询 SysOperationRecord
		operationRecordRouterWithoutRecord.GET("getSysOperationRecordList", operationRecordApi.GetSysOperationRecordList) // 获取 SysOperationRecord 列表
	}
}

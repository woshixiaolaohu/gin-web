package system

import (
	v1 "gin-vue-admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type SysExportTemplateRouter struct{}

func (s *SysExportTemplateRouter) InitSysExportTemplateRouter(Router *gin.RouterGroup) {
	sysExportTemplateRouter := Router.Group("sysExportTemplate").Use(middleware.OperationRecord())
	sysExportTemplateRouterWithoutRecord := Router.Group("sysExportTemplate")
	var sysExportTemplateApi = v1.ApiGroupApp.SystemApiGroup.SysExportTemplateApi
	{
		sysExportTemplateRouter.POST("createSysExportTemplate", sysExportTemplateApi.CreateSysExportTemplate)             // 新建导出模板
		sysExportTemplateRouter.DELETE("deleteSysExportTemplate", sysExportTemplateApi.DeleteSysExportTemplate)           // 删除导出模板
		sysExportTemplateRouter.DELETE("deleteSysExportTemplateByIds", sysExportTemplateApi.DeleteSysExportTemplateByIds) // 批量删除导出模板
		sysExportTemplateRouter.PUT("updateSysExportTemplate", sysExportTemplateApi.UpdateSysExportTemplate)              // 更新导出模板
		sysExportTemplateRouter.POST("importExcel", sysExportTemplateApi.ImportExcel)                                     // 导入文件
	}
	{
		sysExportTemplateRouterWithoutRecord.GET("getSysExportTemplate", sysExportTemplateApi.GetSysExportTemplate)         // 根据 ID 获取导出模板
		sysExportTemplateRouterWithoutRecord.GET("getSysExportTemplateList", sysExportTemplateApi.GetSysExportTemplateList) // 获取导出模板列表
		sysExportTemplateRouterWithoutRecord.GET("exportTemplate", sysExportTemplateApi.ExportTemplate)                     // 导出表格
		sysExportTemplateRouterWithoutRecord.GET("exportTemplate", sysExportTemplateApi.ExportTemplate)                     // 导出模板
	}
}

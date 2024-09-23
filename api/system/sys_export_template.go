package system

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/common/response"
	"gin-vue-admin/model/system"
	systemReq "gin-vue-admin/model/system/request"
	"gin-vue-admin/service"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type SysExportTemplateApi struct{}

var sysExportTemplateService = service.ServiceGroupApp.SystemServiceGroup.SysExportTemplateService

// CreateSysExportTemplate 创建导出模板
// @Tags SysExportTemplate
// @Summary 创建导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "创建导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /sysExportTemplate/createSysExportTemplate [post]
func (s *SysExportTemplateApi) CreateSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	if err := utils.Verify(sysExportTemplate, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := sysExportTemplateService.CreateSysExportTemplate(&sysExportTemplate); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSysExportTemplate 删除导出模板
// @Tags SysExportTemplate
// @Summary 删除导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "删除导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /sysExportTemplate/deleteSysExportTemplate [delete]
func (s *SysExportTemplateApi) DeleteSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage("删除失败", c)
		return
	}
	err = sysExportTemplateService.DeleteSysExportTemplate(sysExportTemplate)
	if err != nil {
		global.GVA_LOG.Error("删除失败", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSysExportTemplateByIds 批量删除导出模板
// @Tags SysExportTemplate
// @Summary 批量删除导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /sysExportTemplate/deleteSysExportTemplateByIds [delete]
func (s *SysExportTemplateApi) DeleteSysExportTemplateByIds(c *gin.Context) {
	var Ids request.IdsReq
	err := c.ShouldBindJSON(&Ids)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysExportTemplateService.DeleteSysExportTemplateByIds(Ids)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSysExportTemplate 更新导出模板
// @Tags SysExportTemplate
// @Summary 更新导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.SysExportTemplate true "更新导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /sysExportTemplate/updateSysExportTemplate [put]
func (s *SysExportTemplateApi) UpdateSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(&sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name": {utils.NotEmpty()},
	}
	err = utils.Verify(sysExportTemplate, verify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysExportTemplateService.UpdateSysExportTemplate(sysExportTemplate)
	if err != nil {
		global.GVA_LOG.Error("更新失败", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetSysExportTemplate 用id查询导出模板
// @Tags SysExportTemplate
// @Summary 用id查询导出模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.SysExportTemplate true "用id查询导出模板"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /sysExportTemplate/getSysExportTemplate [get]
func (s *SysExportTemplateApi) GetSysExportTemplate(c *gin.Context) {
	var sysExportTemplate system.SysExportTemplate
	err := c.ShouldBindJSON(sysExportTemplate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	resysExportTemplate, err := sysExportTemplateService.GetSysExportTemplate(sysExportTemplate.ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败", zap.Error(err))
		response.FailWithMessage("查询失败", c)
		return
	}
	response.OkWithData(gin.H{"resysExportTemplate": resysExportTemplate}, c)
}

// GetSysExportTemplateList 分页获取导出模板列表
// @Tags SysExportTemplate
// @Summary 分页获取导出模板列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query systemReq.SysExportTemplateSearch true "分页获取导出模板列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /sysExportTemplate/getSysExportTemplateList [get]
func (s *SysExportTemplateApi) GetSysExportTemplateList(c *gin.Context) {
	var pageInfo systemReq.SysExportTemplateSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sysExportTemplateService.GetSysTemplateInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// ExportExcel 导出表格
// @Tags SysExportTemplate
// @Summary 导出表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/exportExcel [get]
func (s *SysExportTemplateApi) ExportExcel(c *gin.Context) {
	templateID := c.Query("templateID")
	queryParams := c.Request.URL.Query()
	if templateID == "" {
		response.FailWithMessage("模板ID不能为空", c)
		return
	}
	file, name, err := sysExportTemplateService.ExportExcel(templateID, queryParams)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name+utils.RandomString(6)+".xlsx")) // 对下载的文件重命名
	c.Header("success", "true")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", file.Bytes())
}

// ExportTemplate 导出表格模板
// @Tags SysExportTemplate
// @Summary 导出表格模板
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/exportExcel [get]
func (s *SysExportTemplateApi) ExportTemplate(c *gin.Context) {
	templateID := c.Query("templateID")
	if templateID == "" {
		response.FailWithMessage("模板ID不能为空", c)
		return
	}
	file, name, err := sysExportTemplateService.ExportTemplate(templateID)
	if err != nil {
		global.GVA_LOG.Error("获取失败", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", name+"模板.xslx")) // 下载文件重命名
	c.Header("success", "true")
	c.Data(http.StatusOK, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", file.Bytes())
}

// ImportExcel 导入表格
// @Tags SysImportTemplate
// @Summary 导入表格
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Router /sysExportTemplate/importExcel [post]
func (s *SysExportTemplateApi) ImportExcel(c *gin.Context) {
	templateID := c.Query("templateID")
	if templateID == "" {
		response.FailWithMessage("模板ID不能为空", c)
		return
	}
	file, err := c.FormFile("file")
	if err != nil {
		global.GVA_LOG.Error("文件获取失败", zap.Error(err))
		response.FailWithMessage("文件获取失败", c)
		return
	}
	err = sysExportTemplateService.ImportExcel(templateID, file)
	if err != nil {
		global.GVA_LOG.Error(err.Error(), zap.Error(err))
		response.FailWithMessage(err.Error(), c)
	}
	response.OkWithMessage("导入成功", c)
}

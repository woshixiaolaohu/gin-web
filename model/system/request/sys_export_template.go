package request

import (
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
	"time"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" from:"endCreatedAt"`
	request.PageInfo
}

package request

import (
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
	"time"
)

type SysExportTemplateSearch struct {
	system.SysExportTemplate
	StartCreatedAt *time.Time `json:"start_created_at" form:"start_created_at"`
	EndCreatedAt   *time.Time `json:"end_created_at" from:"end_created_at"`
	request.PageInfo
}

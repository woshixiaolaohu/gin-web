package example

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
)

type ExaCustomer struct {
	global.GVA_MODEL
	CustomerName       string         `json:"customerName"`                               // 客户名称
	CustomerPhone      string         `json:"customerPhone"`                              // 客户联系方式
	SysUserID          uint           `json:"sysUserId"`                                  // 管理ID
	SysUserAuthorityID uint           `json:"sysUserAuthorityId"`                         // 管理角色ID
	SysUser            system.SysUser `json:"sysUser" gorm:"comment:管理详情" form:"sysUser"` //管理详情
}

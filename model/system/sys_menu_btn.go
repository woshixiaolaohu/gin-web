package system

import "gin-vue-admin/global"

type SysBaseMenuBtn struct {
	global.GVA_MODEL
	Name          string `json:"name" gorm:"comment:按钮关键key"`
	Desc          string `json:"desc" gorm:"comment:按钮备注"`
	SysBaseMenuID uint   `json:"sysBaseMenuId" gorm:"comment:菜单ID"`
}

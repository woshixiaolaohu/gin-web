package request

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
)

type AddMenuAuthorityInfo struct {
	Menus       []system.SysBaseMenu `json:"menus"`
	AuthorityId uint                 `json:"authorityId"` // 角色 ID
}

func DefaultMenu() []system.SysBaseMenu {
	return []system.SysBaseMenu{{
		GVA_MODEL: global.GVA_MODEL{ID: 1},
		ParentID:  0,
		Path:      "dashboard",
		Name:      "dashboard",
		Component: "view/dashboard/index.vue",
		Sort:      1,
		Meta: system.Meta{
			Title: "仪表盘",
			Icon:  "setting",
		},
	}}
}

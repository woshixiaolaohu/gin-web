package system

import "gin-vue-admin/global"

type SysApi struct {
	global.GVA_MODEL
	Path        string `json:"path" gorm:"comment:api路径"`             //api路径
	Description string `json:"description" gorm:"comment:api中文描述"`    //api中文描述
	ApiGroup    string `json:"api_group" gorm:"comment:api组"`         //api组
	Method      string `json:"method" gorm:"default:POST;comment:方法"` //方法：创建POST(默认) | GET(查看) | PUT(更新) | DELETE(删除)
}

func (SysApi) TableName() string {
	return "sys_apis"
}

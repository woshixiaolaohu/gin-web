package system

import "gin-vue-admin/global"

type SysDictionary struct {
	global.GVA_MODEL
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:字典名(中文)"`  // 字典名(中文)
	Type                 string                `json:"type" form:"type" gorm:"column:type;comment:字典名(英文)"`  // 字典名(英文)
	Status               *bool                 `json:"status" form:"status" gorm:"column:status;comment:状态"` // 状态
	Desc                 string                `json:"desc" form:"desc" gorm:"column:desc;comment:描述"`       // 描述
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sys_dictionary_details"`
}

func (SysDictionary) TableName() string {
	return "sys_dictionaries"
}

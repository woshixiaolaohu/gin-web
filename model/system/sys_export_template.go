package system

import "gin-vue-admin/global"

// SysExportTemplate 导出模板结构体
type SysExportTemplate struct {
	global.GVA_MODEL
	DBName       string         `json:"db_name" form:"db_name" gorm:"column:db_name;comment:数据库名称"`                                       // 数据库名称
	Name         string         `json:"name" form:"name" gorm:"column:name;comment:模板名称"`                                                 // 模板名称
	TableName    string         `json:"table_name" form:"table_name" gorm:"column:table_name;comment:表名"`                                 // 表名
	TemplateID   string         `json:"template_id" form:"template_id" gorm:"column:template_id;comment:模板标识"`                            //模板标识
	TemplateInfo string         `json:"template_info" form:"template_info" gorm:"column:template_info;comment:模板信息"`                      // 模板信息
	Limit        int            `json:"limit" form:"limit" gorm:"column:limit;comment:导出限制"`                                              // 导出限制
	Order        string         `json:"order" form:"order" gorm:"column:order;comment:排序"`                                                // 排序
	Conditions   []Condition    `json:"conditions" form:"conditions" gorm:"foreignKey:TemplateID;references:TemplateID;comment:条件"`       // 条件
	JoinTemplate []JoinTemplate `json:"join_template" form:"join_template" gorm:"foreignKey:TemplateID;references:TemplateID;comment:关联"` // 关联
}

type Condition struct {
	global.GVA_MODEL
	TemplateID string `json:"template_id" form:"template_id" gorm:"column:template_id;comment:模板标识"` // 模板标识
	Form       string `json:"form" form:"form" gorm:"column:form;comment:条件取得key"`                   // 条件取得key
	Column     string `json:"column" form:"column" gorm:"column:column;comment:作为查询条件的字段"`           // 作为查询条件的字段
	Operator   string `json:"operator" form:"operator" gorm:"column:operator;comment:关联条件"`          // 关联条件
}

type JoinTemplate struct {
	global.GVA_MODEL
	TemplateID string `json:"template_id" form:"template_id" gorm:"column:template_id;comment:模板标识"` // 模板标识
	JOINS      string `json:"joins" form:"joins" gorm:"column:joins;comment:关联"`                     // 关联
	Table      string `json:"table" form:"table" gorm:"column:table;comment:关联表"`                    // 关联表
	ON         string `json:"on" from:"on" gorm:"column:on;comment:关联条件"`                            //关联条件
}

func (Condition) TableName() string {
	return "sys_export_template_conditions"
}

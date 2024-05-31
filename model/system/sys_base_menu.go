package system

import "gin-vue-admin/global"

type SysBaseMenu struct {
	global.GVA_MODEL
	MenuLevel     uint                                       `json:"-"`
	ParentID      uint                                       `json:"parent_id" gorm:"comment:父菜单ID"`    //父菜单ID
	Path          string                                     `json:"path" gorm:"comment:路由path"`        //路由path
	Name          string                                     `json:"name" gorm:"comment:路由name"`        //路由name
	Hidden        bool                                       `json:"hidden" gorm:"comment:是否在列表隐藏"`     //是否在列表隐藏
	Component     string                                     `json:"component" gorm:"comment:对应前端文件路径"` //对应前端文件路径
	Sort          int                                        `json:"sort" gorm:"comment:排序标记"`          //排序标记
	Meta          `json:"meta" gorm:"embedded;comment:附加属性"` //附加属性
	SysAuthoritys []SysAuthority                             `json:"sys_authoritys" gorm:"many2many:sys_authority_menus"`
	Children      []SysBaseMenu                              `json:"children" gorm:"-"`
	Parameters    []SysBaseMenuParameter                     `json:"parameters"`
	MenuBtn       []SysBaseMenuBtn                           `json:"menu_btn"`
}

type Meta struct {
	ActiveName  string `json:"active_name" gorm:"comment:高亮菜单"`     //高亮菜单
	KeepAlive   bool   `json:"keep_alive" gorm:"comment:是否缓存"`      //是否缓存
	DefaultMenu bool   `json:"default_menu" gorm:"comment:是否是基础路由"` //是否是基础路由
	Title       string `json:"title" gorm:"comment:菜单名"`            //菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`            //菜单图标
	CloseTab    bool   `json:"close_tab" gorm:"comment:是否自动关闭tab"`  //是否自动关闭tab
}
type SysBaseMenuParameter struct {
	global.GVA_MODEL
	SysBaseMenuID uint
	Type          string `json:"type" gorm:"comment:地址栏携带参数为params|query"` //地址栏携带参数为params|query
	Key           string `json:"key" gorm:"comment:地址栏携带参数的key"`           //地址栏携带参数的key
	Value         string `json:"value" gorm:"comment:地址栏携带参数的值"`           //地址栏携带参数的值
}

func (SysBaseMenu) TableName() string {
	return "sys_base_menus"
}

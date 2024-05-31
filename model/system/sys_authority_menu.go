package system

type SysMenu struct {
	SysBaseMenu
	MenuID      uint                   `json:"menu_id" gorm:"comment:菜单ID"`      //菜单ID
	AuthorityID uint                   `json:"authority_id" gorm:"comment:角色ID"` //角色ID
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuID"`
	Btns        map[string]uint        `json:"btns" gorm:"-"`
}
type SysAuthorityMenu struct {
	MenuID      string `json:"menu_id" gorm:"comment:菜单ID;column:sys_base_menu_id"`
	AuthorityID string `json:"-" gorm:"comment:角色ID;column:sys_authority_authority_id"`
}

func (SysAuthorityMenu) TableName() string {
	return "sys_authority_menus"
}

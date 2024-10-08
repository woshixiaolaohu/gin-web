package system

import "time"

type SysAuthority struct {
	CreatedAt       time.Time       //创建时间
	UpdatedAt       time.Time       //更新时间
	DeleteAt        *time.Time      `sql:"index"`                                                               //删除时间
	AuthorityId     uint            `json:"authorityId" gorm:"not null;unique;primaryKey;comment:角色ID;size:90"` //角色ID
	AuthorityName   string          `json:"authorityName" gorm:"comment:角色名称"`                                  // 角色名
	ParentID        *uint           `json:"parentID" gorm:"comment:父级角色ID"`                                     //父级角色ID
	DataAuthorityID []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id"`
	Children        []SysAuthority  `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []SysUser       `json:"users" gorm:"many2many:sys_user_authority"`
	DefaultRouter   string          `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (SysAuthority) TableName() string {
	return "sys_authorities"
}

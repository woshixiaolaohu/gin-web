package system

import (
	"gin-vue-admin/global"
	"github.com/gofrs/uuid/v5"
)

type SysUser struct {
	global.GvaModel
	UUID        uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`                                                                       //用户UUID
	UserName    string         `json:"user_name" gorm:"index;comment:用户登录名称"`                                                                  //用户登录名称
	PassWord    string         `json:"-" gorm:"comment:用户登录密码"`                                                                                // 用户登录密码
	NickName    string         `json:"nick_name" gorm:"default:大马猴;comment:用户昵称"`                                                              //用户昵称
	SideMode    string         `json:"side_mode" gorm:"default:dark;comment:用户侧边栏主题"`                                                          //用户侧边栏主题
	Avatar      string         `json:"avatar" gorm:"default:https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png;comment:用户头像"` //用户头像
	BaseColor   string         `json:"base_color" gorm:"default:#fff;comment:基础颜色"`                                                            //基础颜色
	AuthorityID uint           `json:"authority_id" gorm:"default:666666;comment:用户角色ID"`                                                      //用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityID;references:AuthorityID;comment:用户角色"`                            //用户角色
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority"`
	Phone       string         `json:"phone" gorm:"comment:用户手机号"`                      //用户手机号
	Email       string         `json:"email" gorm:"comment:用户邮箱"`                       //用户邮箱
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"` //用户状态

}

func (SysUser) TableName() string {
	return "sys_users"
}

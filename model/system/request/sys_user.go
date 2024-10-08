package request

import "gin-vue-admin/model/system"

type Register struct {
	Username     string `json:"username" example:"用户名"`
	Password     string `json:"password" example:"密码"`
	NickName     string `json:"nickName" example:"昵称"`
	HeaderImg    string `json:"headerImg" example:"头像链接"`
	AuthorityId  uint   `json:"authorityId" swaggertype:"string" example:"uint 角色id"`
	Enable       int    `json:"enable" swaggertype:"string" example:"int 是否启用"`
	AuthorityIds []uint `json:"authorityIds" swaggertype:"string" example:"[]uint 角色id"`
	Phone        string `json:"phone" example:"电话号码"`
	Email        string `json:"email" example:"电子邮箱"`
}

// Login
// user login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
	Password  string `json:"password"`  //密码
	Captcha   string `json:"captcha"`   //验证码
	CaptchaID string `json:"captchaID"` //验证码 ID
}

// ChangePasswordReq
// modify password structure
type ChangePasswordReq struct {
	ID          uint   `json:"-"`           // 从 JWT 中提取 user.id 避免越权
	Password    string `json:"password"`    // 旧密码
	NewPassword string `json:"newPassword"` // 新密码
}

// SetUserAuth
// modify  user's auth structure
type SetUserAuth struct {
	AuthorityId uint `json:"authorityId"` // 角色 ID
}

// SetUserAuthorities
// modify  user's auth structure
type SetUserAuthorities struct {
	ID           uint
	AuthorityIds []uint `json:"authorityIds"` // 角色 ID
}

type ChangeUserInfo struct {
	ID           uint                  `gorm:"primarykey"`                                // 主键 ID
	NickName     string                `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` //用户昵称
	Phone        string                `json:"phone" gorm:"comment:用户手机号"`                // 用户手机号
	AuthorityIds []uint                `json:"authorityIds" gorm:"-"`                     // 角色 ID
	Email        string                `json:"email" gorm:"comment:用户邮箱"`                 // 用户邮箱
	HeaderImg    string                `json:"headerImg"  gorm:"default:https://cube.elemecdn.com/3/7c/3ea6beec64369c2642b92c6726f1epng.png;comment:用户头像"`
	SideMode     string                `json:"sideMode" gorm:"default:dark;comment:用户侧边栏主题"` //用户侧边栏主题
	Enable       int                   `json:"enable" gorm:"comment:冻结用户"`                   // 冻结用户
	Authorities  []system.SysAuthority `json:"-" gorm:"many2many:sys_user_authority;"`
}

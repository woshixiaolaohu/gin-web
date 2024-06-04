package system

// SysUserAuthority sysUser 和 sysAuthority 的连接表
type SysUserAuthority struct {
	SysUserID               uint `gorm:"column:sys_user_id"`
	SysAuthorityAuthorityID uint `gorm:"column:sys_authority_authority_id"`
}

func (s *SysUserAuthority) TableName() string {
	return "sys_user_authorities"
}

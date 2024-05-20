package example

type ExaCustomer struct {
	CustomerName       string `json:"customer_name"`         // 客户名称
	CustomerPhone      string `json:"customer_phone"`        // 客户联系方式
	SysUserID          uint   `json:"sys_user_id"`           // 管理ID
	SysUserAuthorityID uint   `json:"sys_user_authority_id"` // 管理角色ID
}

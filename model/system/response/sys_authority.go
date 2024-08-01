package response

import "gin-vue-admin/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityID uint                `json:"old_authority_id"` // 旧角色 ID
}

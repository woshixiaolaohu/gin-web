package request

type SysAuthorityBtnReq struct {
	MenuID      uint   `json:"menuId"`
	AuthorityId uint   `json:"authorityId"`
	Selected    []uint `json:"selected"`
}

package request

type SysAuthorityBtnReq struct {
	MenuID      uint   `json:"menu_id"`
	AuthorityID uint   `json:"authority_id"`
	Selected    []uint `json:"selected"`
}

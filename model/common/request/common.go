package request

type PageInfo struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"page_size" form:"page_size"`
	KeyWord  string `json:"key_word" form:"key_word"`
}

type GetById struct {
	ID int `json:"id" form:"id"` //主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

type GetAuthorityId struct {
	AuthorityId uint `json:"authority_id" form:"authority_id"` //角色ID
}
type Empty struct {
}

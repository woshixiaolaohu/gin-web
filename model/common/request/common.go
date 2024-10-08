package request

type PageInfo struct {
	Page     int    `json:"page" form:"page"`
	PageSize int    `json:"pageSize" form:"pageSize"`
	KeyWord  string `json:"keyword" form:"keyword"`
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
	AuthorityId uint `json:"authorityId" form:"authorityId"` //角色ID
}
type Empty struct {
}

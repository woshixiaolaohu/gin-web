package request

import (
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
)

// SearchApiParams
// api 分页传条件及排序结构体
type SearchApiParams struct {
	system.SysApi
	request.PageInfo
	OrderKey string `json:"orderKey"` // 排序
	Desc     bool   `json:"desc"`     // 排序方式
}

package system

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
	"strconv"
	"strings"
)

// SysAutoCodeHistory 自动迁移代码记录，用于回滚，重放使用
// omitempty 表示如果该字段的值为空值，即零值或者空引用，则忽略该字段，不将其包含在生成的 json 中
type SysAutoCodeHistory struct {
	global.GVA_MODEL
	Package       string `json:"package"`
	BusinessDB    string `json:"businessDb"`
	TableName     string `json:"tableName"`
	MenuID        uint   `json:"menuID"`
	RequestMeta   string `json:"requestMeta,omitempty" gorm:"type:text"`    // 前端传入结构化信息
	AutoCodePath  string `json:"autoCodePath,omitempty" gorm:"type:text"`   // 其他的meat信息
	InjectionMeta string `json:"injection_meta,omitempty" gorm:"type:text"` // 注入的内容 RouterPath@functionName@RouterString
	StructName    string `json:"structName"`
	StructCNName  string `json:"structCNName"`
	ApiIDs        string `json:"apiIds,omitempty"` // api表注册内容
	Flag          int    `json:"flag"`             // 表示对应状态 0 代表创建 1代表回滚
}

// ToRequestIds ApiIDs 转换 request.IdsReq
func (s *SysAutoCodeHistory) ToRequestIds() request.IdsReq {
	if s.ApiIDs == "" {
		return request.IdsReq{}
	}
	slice := strings.Split(s.ApiIDs, ";")
	ids := make([]int, 0, len(slice))
	length := len(slice)
	for i := 0; i < length; i++ {
		id, _ := strconv.ParseInt(slice[i], 10, 32)
		ids = append(ids, int(id))
	}
	return request.IdsReq{
		Ids: ids,
	}
}

package system

import "gin-vue-admin/config"

// System 配置文件结构体
type System struct {
	Config config.Server `json:"config"`
}

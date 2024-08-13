package api

import (
	"gin-vue-admin/api/example"
	"gin-vue-admin/api/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

package router

import (
	"gin-vue-admin/router/example"
	"gin-vue-admin/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

package system

import "gin-vue-admin/service"

type ApiGroup struct {
	BaseApi
	AuthorityApi
	AuthorityBtnApi
	SystemApiApi
	CasbinApi
	DictionaryApi
	DictionaryDetailApi
	SysExportTemplateApi
	DBApi
	JwtApi
	AuthorityMenuApi
}

var (
	apiService              = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService              = service.ServiceGroupApp.SystemServiceGroup.JwtService
	menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	baseMenuService         = service.ServiceGroupApp.SystemServiceGroup.BaseMenuService
	userService             = service.ServiceGroupApp.SystemServiceGroup.UserService
	casbinService           = service.ServiceGroupApp.SystemServiceGroup.CasbinService
	authorityService        = service.ServiceGroupApp.SystemServiceGroup.AuthorityService
	authorityBtnService     = service.ServiceGroupApp.SystemServiceGroup.AuthorityBtnService
	dictionaryService       = service.ServiceGroupApp.SystemServiceGroup.DictionaryService
	dictionaryDetailService = service.ServiceGroupApp.SystemServiceGroup.DictionaryDetailService
	initDBService           = service.ServiceGroupApp.SystemServiceGroup.InitDBService
)

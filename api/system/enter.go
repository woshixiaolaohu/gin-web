package system

import "gin-vue-admin/service"

type ApiGroup struct {
	BaseApi
}

var (
	apiService    = service.ServiceGroupApp.SystemServiceGroup.ApiService
	jwtService    = service.ServiceGroupApp.SystemServiceGroup.JwtService
	userService   = service.ServiceGroupApp.SystemServiceGroup.UserService
	casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
)

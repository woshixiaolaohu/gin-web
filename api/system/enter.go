package system

import "gin-vue-admin/service"

type ApiGroup struct {
	BaseApi
}

var (
	apiService    = service.ServiceGroupApp.SystemServiceGroup.ApiService
	casbinService = service.ServiceGroupApp.SystemServiceGroup.CasbinService
)

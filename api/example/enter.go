package example

import "gin-vue-admin/service"

type ApiGroup struct {
	CustomerApi
}

var customerService = service.ServiceGroupApp.ExampleServiceGroup.CustomerService

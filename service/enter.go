package service

import "gin-vue-admin/service/example"

type ServiceGroup struct {
	ExampleServiceGroup example.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)

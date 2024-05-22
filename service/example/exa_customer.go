package example

import "gin-vue-admin/model/example"
import "gin-vue-admin/global"

type CustomerService struct{}

func (exa *CustomerService) CreateExaCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.create()
}

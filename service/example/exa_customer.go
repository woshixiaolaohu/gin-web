package example

import (
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/example"
	"gin-vue-admin/model/system"
	systemService "gin-vue-admin/service/system"
)
import "gin-vue-admin/global"

type CustomerService struct{}

// CreateExaCustomer
// @function: CreateExaCustomer
// @description: 创建客户
// @param: e model.ExaCustomer
// @return: err error
func (exa *CustomerService) CreateExaCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

// DeleteExaCustomer
// @function: DeleteFileChunk
// @description: 删除客户
// @param: e model.ExaCustomer
// @return: err error
func (exa *CustomerService) DeleteExaCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.Delete(&e).Error
	return err
}

// UpdateCustomer
// @function: UpdateExaCustomer
// @description: 更新客户
// @param: e *model.ExaCustomer
// @return: err error
func (exa *CustomerService) UpdateCustomer(e example.ExaCustomer) (err error) {
	err = global.GVA_DB.Save(&e).Error
	return err
}

// GetExaCustomer
// @function: GetExaCustomer
// @description: 获取客户信息
// @param: id uint
// @return: customer model.ExaCustomer, err error
func (exa *CustomerService) GetExaCustomer(id uint) (customer example.ExaCustomer, err error) {
	err = global.GVA_DB.Where("id = ? ", id).First(&customer).Error
	return
}

func (exa *CustomerService) GetCustomerInfoList(sysUserAuthorityID uint, info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.PageSize - 1)
	db := global.GVA_DB.Model(&example.ExaCustomer{})
	var a system.SysAuthority
	a.AuthorityID = sysUserAuthorityID
	auth, err := systemService.
}

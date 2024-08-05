package example

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/example"
	"gin-vue-admin/model/system"
	systemService "gin-vue-admin/service/system"
)

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

// UpdateExaCustomer
// @function: UpdateExaCustomer
// @description: 更新客户
// @param: e *model.ExaCustomer
// @return: err error
func (exa *CustomerService) UpdateExaCustomer(e *example.ExaCustomer) (err error) {
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
	auth, err := systemService.AuthorityServiceApp.GetAuthorityInfo(a)
	if err != nil {
		return
	}
	var dataID []uint
	for _, v := range auth.DataAuthorityID {
		dataID = append(dataID, v.AuthorityID)
	}
	var CustomerList []example.ExaCustomer
	err = db.Where("sys_user_authority_id in ?", dataID).Count(&total).Error
	if err != nil {
		return CustomerList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataID).Find(&CustomerList).Error
	}
	return CustomerList, total, err
}

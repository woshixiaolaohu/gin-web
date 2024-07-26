package system

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
)

var ErrRoleExistence = errors.New("存在相同角色id")

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

// CreateAuthority
// @function: CreateAuthority
// @description: 创建一个角色
// @param: auth model.SysAuthority
// @return: authority system.SysAuthority, err error
func (a *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	if err = global.GVA_DB.Where("") {

	}
}

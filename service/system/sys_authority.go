package system

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
	systemReq "gin-vue-admin/model/system/request"
	"gin-vue-admin/model/system/response"
	"gorm.io/gorm"
	"strconv"
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
func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	if err = global.GVA_DB.Where("authority_id= ?", auth.AuthorityID).First(&system.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		return auth, ErrRoleExistence
	}

	e := global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Create(&auth).Error; err != nil {
			return err
		}
		auth.SysBaseMenus = systemReq.DefaultMenu()
		if err = tx.Model(&auth).Association("SysBaseMenus").Replace(&auth.SysBaseMenus); err != nil {
			return err
		}
		casbinInfos := systemReq.DefaultCasbin()
		authorityID := strconv.Itoa(int(auth.AuthorityID))
		rules := [][]string{}
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityID, v.Path, v.Method})
		}
		return CasbinServiceApp.AddPolicies(tx, rules)
	})
	return auth, e
}

// CopyAuthority
// @function: CopyAuthority
// @description: 复制一个角色
// @param: copyInfo response.SysAuthorityCopyResponse
// @return: authority system.SysAuthority, err error
func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ? ", copyInfo.Authority.AuthorityID).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return authority, ErrRoleExistence
	}
	copyInfo.Authority.Children = []system.SysAuthority{}
}

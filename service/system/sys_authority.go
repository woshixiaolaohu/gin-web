package system

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
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
//
//	@function:		CreateAuthority
//	@description:	创建一个角色
//	@param:			auth model.SysAuthority
//	@return:		authority system.SysAuthority, err error
func (authorityService *AuthorityService) CreateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	if err = global.GVA_DB.Where("authority_id= ?", auth.AuthorityId).First(&system.SysAuthority{}).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
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
		authorityID := strconv.Itoa(int(auth.AuthorityId))
		rules := [][]string{}
		for _, v := range casbinInfos {
			rules = append(rules, []string{authorityID, v.Path, v.Method})
		}
		return CasbinServiceApp.AddPolicies(tx, rules)
	})
	return auth, e
}

// CopyAuthority
//
//	@function:		CopyAuthority
//	@description:	复制一个角色
//	@param:			copyInfo response.SysAuthorityCopyResponse
//	@return:		authority system.SysAuthority, err error
func (authorityService *AuthorityService) CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (authority system.SysAuthority, err error) {
	var authorityBox system.SysAuthority
	if !errors.Is(global.GVA_DB.Where("authority_id = ? ", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return authority, ErrRoleExistence
	}
	copyInfo.Authority.Children = []system.SysAuthority{}
	menus, err := MenuServiceApp.GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityID})
	if err != nil {
		return
	}
	var baseMenu []system.SysBaseMenu
	for _, v := range menus {
		intNum := v.MenuID
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = global.GVA_DB.Create(&copyInfo.Authority).Error
	if err != nil {
		return
	}
	var btns []system.SysAuthorityBtn
	err = global.GVA_DB.Find(&btns, "authority_id = ?", copyInfo.OldAuthorityID).Error
	if err != nil {
		return
	}
	if len(btns) > 0 {
		for i := range btns {
			btns[i].AuthorityId = copyInfo.Authority.AuthorityId
		}
		err = global.GVA_DB.Create(&btns).Error
		if err != nil {
			return
		}
	}
	paths := CasbinServiceApp.GetPolicyPathByAuthorityID(copyInfo.OldAuthorityID)
	err = CasbinServiceApp.UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = authorityService.DeleteAuthority(&copyInfo.Authority)
	}
	return copyInfo.Authority, err
}

// UpdateAuthority
//
//	@function:		UpdateAuthority
//	@description:	更改一个角色
//	@param:			auth model.SysAuthority
//	@return:		authority system.SysAuthority, err error
func (authorityService *AuthorityService) UpdateAuthority(auth system.SysAuthority) (authority system.SysAuthority, err error) {
	var oldAuthority system.SysAuthority
	err = global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&oldAuthority).Error
	if err != nil {
		global.GVA_LOG.Debug(err.Error())
		return system.SysAuthority{}, errors.New("查询角色数据失败")
	}
	err = global.GVA_DB.Model(&oldAuthority).Updates(&auth).Error
	return auth, err
}

// DeleteAuthority
//
//	@function:		DeleteAuthority
//	@description:	删除角色
//	@param:			auth *model.SysAuthority
//	@return:		err error
func (authorityService *AuthorityService) DeleteAuthority(auth *system.SysAuthority) error {
	if errors.Is(global.GVA_DB.Debug().Preload("Users").First(&auth).Error, gorm.ErrRecordNotFound) {
		return errors.New("该角色不存在")
	}
	if len(auth.Users) != 0 {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("authority_id = ?", auth.AuthorityId).First(&system.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用，禁止删除")
	}
	if !errors.Is(global.GVA_DB.Where("parent_id = ?", auth.ParentID).First(&system.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色，不允许删除")
	}
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Preload("SysBaseMenus").Preload("DataAuthorityID").Where("authority_id = ?", auth.AuthorityId).First(auth).Unscoped().Delete(auth).Error; err != nil {
			return err
		}
		if len(auth.SysBaseMenus) > 0 {
			if err = tx.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus); err != nil {
				return err
			}
		}
		if len(auth.DataAuthorityID) > 0 {
			if err = tx.Model(&auth).Association("DataAuthorityID").Delete(auth.DataAuthorityID); err != nil {
				return err
			}
		}
		if err = tx.Delete(&system.SysUserAuthority{}, "sys_authority_authority_id = ?", auth.AuthorityId).Error; err != nil {
			return err
		}
		if err = tx.Where("authority_id = ?", auth.AuthorityId).Delete(&[]system.SysAuthorityBtn{}).Error; err != nil {
			return err
		}
		authorityID := strconv.Itoa(int(auth.AuthorityId))
		if err = CasbinServiceApp.RemoveFilteredPolicy(tx, authorityID); err != nil {
			return err
		}
		return nil
	})
}

// GetAuthorityInfoList
//
//	@function:		GetAuthorityInfoList
//	@description:	分页获取数据
//	@param:			info request.PageInfo
//	@return:		list interface{}, total int64, err error
func (authorityService *AuthorityService) GetAuthorityInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysAuthority{})
	if err = db.Where("parent_id = ?", "0").Count(&total).Error; total == 0 || err != nil {
		return
	}
	var authority []system.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityID").Where("parent_id = ?", "0").Find(&authority).Error
	for k := range authority {
		err = authorityService.findChildrenAuthority(&authority[k])
	}
	return authority, total, err
}

// GetAuthorityInfo
//
//	@function:		GetAuthorityInfo
//	@description:	获取所有角色信息
//	@param:			auth model.SysAuthority
//	@return:		sa system.SysAuthority, err error
func (authorityService *AuthorityService) GetAuthorityInfo(auth system.SysAuthority) (sa system.SysAuthority, err error) {
	err = global.GVA_DB.Preload("DataAuthorityID").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return sa, err
}

// SetDataAuthority
//
//	@function:		SetDataAuthority
//	@description:	设置角色资源权限
//	@param:			auth model.SysAuthority
//	@return:		error
func (authorityService *AuthorityService) SetDataAuthority(auth system.SysAuthority) error {
	var s system.SysAuthority
	global.GVA_DB.Preload("DataAuthorityID").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("DataAuthorityID").Replace(&auth.SysBaseMenus)
	return err
}

// SetMenuAuthority
//
//	@function:		SetMenuAuthority
//	@description:	菜单与角色绑定
//	@param:			auth *model.SysAuthority
//	@return:		error
func (authorityService *AuthorityService) SetMenuAuthority(auth *system.SysAuthority) error {
	var s system.SysAuthority
	global.GVA_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := global.GVA_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

// findChildrenAuthority
//
//	@function:		findChildrenAuthority
//	@description:	查询子角色
//	@param:			authority *model.SysAuthority
//	@return:		err error
func (authorityService *AuthorityService) findChildrenAuthority(authority *system.SysAuthority) (err error) {
	err = global.GVA_DB.Preload("DataAuthorityID").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = authorityService.findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}

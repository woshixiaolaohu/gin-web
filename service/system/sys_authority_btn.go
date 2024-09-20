package system

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
	"gin-vue-admin/model/system/request"
	"gin-vue-admin/model/system/response"
	"gorm.io/gorm"
)

type AuthorityBtnService struct{}

// GetAuthorityBtn
// @function: GetAuthorityBtn
// @description: 获取btn
// @param: req request.SysAuthorityBtnReq
// @return: res response.SysAuthorityBtnRes, err error
func (authorityBtnService *AuthorityBtnService) GetAuthorityBtn(req request.SysAuthorityBtnReq) (res response.SysAuthorityBtnRes, err error) {
	var authorityBtn []system.SysAuthorityBtn
	err = global.GVA_DB.Find(&authorityBtn, "authority_id = ? AND sys_menu_id = ?", req.AuthorityID, req.MenuID).Error
	if err != nil {
		return
	}
	var selected []uint
	for _, v := range authorityBtn {
		selected = append(selected, v.SysBaseMenuBtnID)
	}
	res.Selected = selected
	return res, err
}

// SetAuthorityBtn
// @function: SetAuthorityBtn
// @description: 设置btn
// @param: req request.SysAuthorityBtnReq
// @return: err error
func (authorityBtnService *AuthorityBtnService) SetAuthorityBtn(req request.SysAuthorityBtnReq) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var authorityBtn []system.SysAuthorityBtn
		err = tx.Delete(&[]system.SysAuthorityBtn{}, "authority_id = ? AND sys_menu_id = ?", req.AuthorityID, req.MenuID).Error
		if err != nil {
			return err
		}
		for _, v := range req.Selected {
			authorityBtn = append(authorityBtn, system.SysAuthorityBtn{
				AuthorityID:      req.AuthorityID,
				SysMenuID:        req.MenuID,
				SysBaseMenuBtnID: v,
			})
		}
		if len(authorityBtn) > 0 {
			err = tx.Create(&authorityBtn).Error
		}
		if err != nil {
			return err
		}
		return err
	})
}

// CanRemoveAuthorityBtn
// @function: CanRemoveAuthorityBtn
// @description: 删除权限按钮
// @param: ID string
// @return: err error
func (authorityBtnService *AuthorityBtnService) CanRemoveAuthorityBtn(ID string) (err error) {
	fErr := global.GVA_DB.First(&system.SysAuthorityBtn{}, "sys_base_menu_btn_id = ?", ID).Error
	if errors.Is(fErr, gorm.ErrRecordNotFound) {
		return nil
	}
	return errors.New("此按钮正在使用无法删除")
}

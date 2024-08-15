package system

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
	"gin-vue-admin/utils"
	"github.com/gofrs/uuid/v5"
	"gorm.io/gorm"
	"time"
)

type UserService struct {
}

// Register
// @function: Register
// @description: 用户注册
// @param: u model.SysUser
// @return: userInter system.SysUser, err error
func (userService *UserService) Register(u system.SysUser) (userInter system.SysUser, err error) {
	var user system.SysUser
	// 判断用户名是否注册
	if !errors.Is(global.GVA_DB.Where("user_name = ?", u.UserName).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户名已注册")
	}
	// 否则 附加 uuid 密码 hash 加密 注册
	u.PassWord = utils.BcryptHash(u.PassWord)
	u.UUID = uuid.Must(uuid.NewV4())
	err = global.GVA_DB.Create(&u).Error
	return u, err
}

// Login
// @function: Login
// @description: 用户登录
// @param: u *model.SysUser
// @return: err error, userInter *model.SysUser
func (userService *UserService) Login(u *system.SysUser) (userInter *system.SysUser, err error) {
	if global.GVA_DB == nil {
		return nil, fmt.Errorf("db not init")
	}
	var user system.SysUser
	err = global.GVA_DB.Where("user_name = ?", u.UserName).Preload("Authorities").Preload("Authority").First(&user).Error
	if err != nil {
		if ok := utils.BcryptCheck(u.PassWord, user.PassWord); !ok {
			return nil, errors.New("密码错误")
		}
		MenuServiceApp.UserAuthorityDefaultRouter(&user)
	}
	return &user, err
}

// ChangePassword
// @function: ChangePassword
// @description: 修改用户密码
// @param: u *model.SysUser, newPassword string
// @return: userInter *model.SysUser,err error
func (userService *UserService) ChangePassword(u *system.SysUser, newPassword string) (userInter *system.SysUser, err error) {
	var user system.SysUser
	if err = global.GVA_DB.Where("id = ?", u.ID).First(&user).Error; err != nil {
		return nil, err
	}
	if ok := utils.BcryptCheck(u.PassWord, user.PassWord); !ok {
		return nil, errors.New("原密码错误")
	}
	user.PassWord = utils.BcryptHash(newPassword)
	err = global.GVA_DB.Save(&user).Error
	return &user, err
}

// GetUserInfoList
// @function: GetUserInfoList
// @description: 分页获取数据
// @param: info request.PageInfo
// @return: err error, list interface{}, total int64
func (userService *UserService) GetUserInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})
	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Preload("Authorities").Preload("Authority").Find(&userList).Error
	return userList, total, err
}

// SetUserAuthority
// @function: SetUserAuthority
// @description: 设置一个用户的权限
// @param: id uint, authorityId string
// @return: err error
func (userService *UserService) SetUserAuthority(id uint, authorityID uint) (err error) {
	assignErr := global.GVA_DB.Where("sys_user_id = ? AND sys_authority_authority_id = ?", id, authorityID).First(&system.SysUserAuthority{}).Error
	if errors.Is(assignErr, gorm.ErrRecordNotFound) {
		return errors.New("该用户无此角色")
	}
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", id).Update("authority_id", authorityID).Error
	return err
}

// SetUserAuthorities
// @function: SetUserAuthorities
// @description: 设置一个用户的权限
// @param: id uint, authorityIds []string
// @return: err error
func (userService *UserService) SetUserAuthorities(id uint, authorityIds []uint) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var user system.SysUser
		TxErr := tx.Where("id = ?", id).First(&user).Error
		if TxErr != nil {
			global.GVA_LOG.Debug(TxErr.Error())
			return errors.New("查询用户数据失败")
		}
		TxErr = tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error
		if TxErr != nil {
			return TxErr
		}
		var userAuthority []system.SysUserAuthority
		for _, v := range authorityIds {
			userAuthority = append(userAuthority, system.SysUserAuthority{
				SysUserID:               id,
				SysAuthorityAuthorityID: v,
			})
		}
		TxErr = tx.Create(&userAuthority).Error
		if TxErr != nil {
			return TxErr
		}
		TxErr = tx.Model(&user).Update("authority_id = ?", authorityIds[0]).Error
		if TxErr != nil {
			return TxErr
		}
		return nil
	})
}

// DeleteUser
// @function: DeleteUser
// @description: 删除用户
// @param: id float64
// @return: err error
func (userService *UserService) DeleteUser(id int) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&system.SysUser{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&[]system.SysUserAuthority{}, "sys_user_id = ?", id).Error; err != nil {
			return err
		}
		return nil
	})
}

// SetUserInfo
// @function: SetUserInfo
// @description: 设置用户信息
// @param: reqUser model.SysUser
// @return: err error, user model.SysUser
func (userService *UserService) SetUserInfo(u system.SysUser) error {
	return global.GVA_DB.Model(&system.SysUser{}).Select("updated_at", "nick_name", "header_img", "phone", "email", "side_mode", "enable").
		Where("id = ?", u.ID).Updates(map[string]interface{}{
		"updated_at": time.Now(),
		"nick_name":  u.NickName,
		"header_img": u.HeaderImg,
		"phone":      u.Phone,
		"email":      u.Email,
		"side_mode":  u.SideMode,
		"enable":     u.Enable,
	}).Error
}

// SetSelfInfo
// @function: SetSelfInfo
// @description: 设置个人信息
// @param: reqUser model.SysUser
// @return: err error, user model.SysUser
func (userService *UserService) SetSelfInfo(u system.SysUser) error {
	return global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", u.ID).Updates(u).Error
}

// GetUserInfo
// @function: GetUserInfo
// @description: 获取用户信息
// @param: uuid uuid.UUID
// @return: err error, user system.SysUser
func (userService *UserService) GetUserInfo(uuid uuid.UUID) (user system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.GVA_DB.Preload("Authorities").Preload("Authority").First(&reqUser, "uuid = ?", uuid).Error
	if err != nil {
		return reqUser, err
	}
	MenuServiceApp.UserAuthorityDefaultRouter(&reqUser)
	return reqUser, err
}

// FindUserById
// @function: FindUserById
// @description: 通过id获取用户信息
// @param: id int
// @return: err error, user *model.SysUser
func (userService *UserService) FindUserById(id int) (user *system.SysUser, err error) {
	var reqUser system.SysUser
	err = global.GVA_DB.Where("id = ?", id).First(&reqUser).Error
	return &reqUser, err
}

// FindUserByUuid
// @function: FindUserByUuid
// @description: 通过uuid获取用户信息
// @param: uuid string
// @return: err error, user *model.SysUser
func (userService *UserService) FindUserByUuid(uuid string) (user *system.SysUser, err error) {
	var reqUser system.SysUser
	if err = global.GVA_DB.Where("uuid = ?", uuid).First(&reqUser).Error; err != nil {
		return &reqUser, errors.New("用户不存在")
	}
	return &reqUser, nil
}

// ResetPassword
// @function: ResetPassword
// @description: 修改用户密码
// @param: ID uint
// @return: err error
func (userService *UserService) ResetPassword(ID uint) (err error) {
	err = global.GVA_DB.Model(&system.SysUser{}).Where("id = ?", ID).Update("password", utils.BcryptHash("123456")).Error
	return err
}

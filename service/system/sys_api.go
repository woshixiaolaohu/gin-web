package system

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
	"gorm.io/gorm"
)

type ApiService struct{}

var ApiServiceApp = new(ApiService)

// CreateApi
//
//	@function:		CreateApi

// @description:	新增基础api
// @param:			api model.SysApi
// @return:		err error
func (apiService *ApiService) CreateApi(api system.SysApi) (err error) {
	if !errors.Is(global.GVA_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&system.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同API")
	}
	return global.GVA_DB.Create(&api).Error
}

// DeleteApi
//
//	@function:		DeleteApi
//	@description:	删除基础api
//	@param:			api model.SysApi
//	@return:		err error
func (apiService *ApiService) DeleteApi(api system.SysApi) (err error) {
	var entity system.SysApi
	err = global.GVA_DB.First(&entity, "id = ?", api.ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = global.GVA_DB.Delete(&entity).Error
	if err != nil {
		return err
	}
	CasbinServiceApp.ClearCasbin(1, entity.Path, entity.Method)
	if err != nil {
		return err
	}
	return nil
}

// GetApiInfoList
//
//	@function:		GetApiInfoList
//	@description:	分页获取数据,
//	@param:			api model.SysApi, info request.PageInfo, order string, desc bool
//	@return:		list interface{}, total int64, err error
func (apiService *ApiService) GetApiInfoList(api system.SysApi, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysApi{})
	var apiList []system.SysApi
	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}
	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}
	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}
	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}
	err = db.Count(&total).Error
	if err != nil {
		return apiList, total, err
	}
	db = db.Limit(limit).Offset(offset)
	OrderStr := "id desc"
	if order != "" {
		orderMap := make(map[string]bool, 5)
		orderMap["id"] = true
		orderMap["path"] = true
		orderMap["api_group"] = true
		orderMap["description"] = true
		orderMap["method"] = true
		if !orderMap[order] {
			err = fmt.Errorf("非法分排序字段: %v", order)
			return apiList, total, err
		}
		OrderStr = order
		if desc {
			OrderStr = order + " desc"
		}
	}
	err = db.Order(OrderStr).Find(&apiList).Error
	return apiList, total, err
}

// GetAllApis
//
//	@function:		GetAllApis
//	@description:	获取所有的api
//	@return:		apis []model.SysApi, err error
func (apiService *ApiService) GetAllApis() (apis []system.SysApi, err error) {
	err = global.GVA_DB.Find(&apis).Error
	return
}

// GetApiByID
//
//	@function:		GetApiByID
//	@description:	根据id获取api
//	@param:			id float64
//	@return:		api model.SysApi, err error
func (apiService *ApiService) GetApiByID(id int) (api system.SysApi, err error) {
	err = global.GVA_DB.First(&api, "id = ?", id).Error
	return
}

// UpdateApi
//
//	@function:		UpdateApi
//	@description:	根据id更新api
//	@param:			api model.SysApi
//	@return:		err error
func (apiService *ApiService) UpdateApi(api system.SysApi) (err error) {
	var oldA system.SysApi
	err = global.GVA_DB.First(&oldA, "id = ?", api.ID).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		var duplaicateApi system.SysApi
		if ferr := global.GVA_DB.First(&duplaicateApi, "path = ? AND method = ?", api.Path, api.Method).Error; ferr != nil {
			if !errors.Is(ferr, gorm.ErrRecordNotFound) {
				return ferr
			}
		} else {
			if duplaicateApi.ID != api.ID {
				return errors.New("存在相同api路径")
			}
		}
	}
	if err != nil {
		return err
	}
	err = CasbinServiceApp.UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
	if err != nil {
		return err
	}
	return global.GVA_DB.Save(&api).Error
}

// DeleteApisByIds
//
//	@function:		DeleteApisByIds
//	@description:	删除选中API
//	@param:			apis []model.SysApi
//	@return:		err error
func (apiService *ApiService) DeleteApisByIds(ids request.IdsReq) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		var apis []system.SysApi
		err = tx.Find(&apis, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		err = tx.Delete(&[]system.SysApi{}, "id in ?", ids.Ids).Error
		if err != nil {
			return err
		}
		for _, sysApi := range apis {
			CasbinServiceApp.ClearCasbin(1, sysApi.Path, sysApi.Method)
		}
		return err
	})
}

package system

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
	"gin-vue-admin/model/system/request"
)

type DictionaryDetailService struct{}

// CreateSysDictionaryDetail
//
//	@function:		CreateSysDictionaryDetail
//	@description:	创建字典详情数据
//	@param:			sysDictionaryDetail model.SysDictionaryDetail
//	@return:		err error
func (dictionaryDetailService *DictionaryDetailService) CreateSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Create(&sysDictionaryDetail).Error
	return err
}

// DeleteSysDictionaryDetail
//
//	@function:		DeleteSysDictionaryDetail
//	@description:	删除字典详情数据
//	@param:			sysDictionaryDetail model.SysDictionaryDetail
//	@return:		err error
func (dictionaryDetailService *DictionaryDetailService) DeleteSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Delete(&sysDictionaryDetail).Error
	return err
}

// UpdateSysDictionaryDetail
//
//	@function:		UpdateSysDictionaryDetail
//	@description:	更新字典详情数据
//	@param:			sysDictionaryDetail *model.SysDictionaryDetail
//	@return:		err error
func (dictionaryDetailService *DictionaryDetailService) UpdateSysDictionaryDetail(sysDictionaryDetail system.SysDictionaryDetail) (err error) {
	err = global.GVA_DB.Save(&sysDictionaryDetail).Error
	return err
}

// GetSysDictionaryDetail
//
//	@function:		GetSysDictionaryDetail
//	@description:	根据id获取字典详情单条数据
//	@param:			id uint
//	@return:		sysDictionaryDetail system.SysDictionaryDetail, err error
func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetail(id uint) (sysDictionaryDetail system.SysDictionaryDetail, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

// GetSysDictionaryDetailInfoList
//
//	@function:		GetSysDictionaryDetailInfoList
//	@description:	分页获取字典详情列表
//	@param:			info request.SysDictionaryDetailSearch
//	@return:		list interface{}, total int64, err error
func (dictionaryDetailService *DictionaryDetailService) GetSysDictionaryDetailInfoList(info request.SysDictionarySearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{})
	var sysDictionaryDetails []system.SysDictionaryDetail
	// 如果有搜索条件 下方自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != "" {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("sort").Find(&sysDictionaryDetails).Error
	return sysDictionaryDetails, total, err
}

// GetDictionaryList
//
//	@function:		GetDictionaryList
//	@description:	按照字典id获取字典全部内容的方法
//	@param:			dictionaryID uint
//	@return:		list []system.SysDictionaryDetail, err error
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryList(dictionaryID uint) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	err = global.GVA_DB.Find(&sysDictionaryDetails, "sys_dictionary_id = ?", dictionaryID).Error
	return sysDictionaryDetails, err
}

// GetDictionaryListByType
//
//	@function:		GetDictionaryListByType
//	@description:	按照字典type获取字典全部内容的方法
//	@param:			dictionaryType string
//	@return:		list []system.SysDictionaryDetail, err error
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryListByType(dictionaryType string) (list []system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails []system.SysDictionaryDetail
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries_id = sys_dictionary_details.sys_dictionary_id")
	err = db.Debug().Find(&sysDictionaryDetails, "type = ?", dictionaryType).Error
	return sysDictionaryDetails, err
}

// GetDictionaryInfoByValue
//
//	@function:		GetDictionaryInfoByValue
//	@description:	按照字典id+字典内容value获取单条字典内容
//	@param:			dictionaryID uint, value string
//	@return:		detail system.SysDictionaryDetail, err error
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryInfoByValue(dictionaryID uint, value string) (detail system.SysDictionaryDetail, err error) {
	var sysDictionaryDetail system.SysDictionaryDetail
	err = global.GVA_DB.First(&sysDictionaryDetail, "sys_dictionary_id = ? AND value = ?", dictionaryID, value).Error
	return sysDictionaryDetail, err
}

// GetDictionaryInfoByTypeValue
//
//	@function:		GetDictionaryInfoByTypeValue
//	@description:	按照字典type+字典内容value获取单条字典内容
//	@param:			dictionaryType string, value string
//	@return:		detail system.SysDictionaryDetail, err error
func (dictionaryDetailService *DictionaryDetailService) GetDictionaryInfoByTypeValue(dictionaryType string, value string) (detail system.SysDictionaryDetail, err error) {
	var sysDictionaryDetails system.SysDictionaryDetail
	db := global.GVA_DB.Model(&system.SysDictionaryDetail{}).Joins("JOIN sys_dictionaries ON sys_dictionaries.id = sys_dictionary_details.sys_dictionary_id")
	err = db.First(&sysDictionaryDetails, "sys_dictionaries.type = ? and sys_dictionary_details.value = ?", dictionaryType, value).Error
	return sysDictionaryDetails, err
}

package system

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
	"gorm.io/gorm"
)

type DictionaryService struct{}

// CreateSysDictionary
//
//	@function:		CreateSysDictionary
//	@description:	创建字典数据
//	@param:			SysDictionary model.SysDictionary
//	@return:		err error
func (dictionaryService *DictionaryService) CreateSysDictionary(sysDictionary system.SysDictionary) (err error) {
	if !(errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同type, 不允许创建")
	}
	err = global.GVA_DB.Create(&sysDictionary).Error
	return err
}

// DeleteSysDictionary
//
//	@function:		DeleteSysDictionary
//	@description:	删除字典数据
//	@param:			sysDictionary model.SysDictionary
//	@return:		err error
func (dictionaryService *DictionaryService) DeleteSysDictionary(sysDictionary system.SysDictionary) (err error) {
	err = global.GVA_DB.Where("id = ?", sysDictionary.ID).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("不允许删除")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&sysDictionary).Error
	if err != nil {
		return err
	}
	if sysDictionary.SysDictionaryDetails != nil {
		return global.GVA_DB.Where("sys_dictionary_id = ?", sysDictionary.ID).Delete(sysDictionary.SysDictionaryDetails).Error
	}
	return
}

// UpdateSysDictionary
//
//	@function:		UpdateSysDictionary
//	@description:	更新字典数据
//	@param:			sysDictionary *model.SysDictionary
//	@return:		err error
func (dictionaryService *DictionaryService) UpdateSysDictionary(sysDictionary *system.SysDictionary) (err error) {
	var dict system.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	err = global.GVA_DB.Where("id = ?", sysDictionary.ID).First(&dict).Error
	if err != nil {
		global.GVA_LOG.Debug(err.Error())
		return errors.New("查询字典数据失败")
	}
	if dict.Type != sysDictionary.Type {
		if errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同type, 不允许创建")
		}
	}
	err = global.GVA_DB.Model(&dict).Updates(sysDictionaryMap).Error
	return err
}

// GetSysDictionary
//
//	@function:		GetSysDictionary
//	@description:	根据id或者type获取字典单条数据
//	@param:			Type string, ID uint
//	@return:		err error, sysDictionary model.SysDictionary
func (dictionaryService *DictionaryService) GetSysDictionary(Type string, ID uint, Status *bool) (sysDictionary system.SysDictionary, err error) {
	var flag = false
	if Status != nil {
		flag = true
	} else {
		flag = *Status
	}
	err = global.GVA_DB.Where("(type = ? OR id = ?) AND status = ?", Type, ID, flag).Preload("SysDictionaryDetails", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", true).Order("sort")
	}).First(&sysDictionary).Error
	return
}

// GetSysDictionaryInfoList
//
//	@function:		GetSysDictionaryInfoList
//	@description:	分页获取字典列表
//	@param:			info request.SysDictionarySearch
//	@return:		err error, list interface{}, total int64
func (dictionaryService *DictionaryService) GetSysDictionaryInfoList() (list interface{}, err error) {
	var sysDictionarys []system.SysDictionary
	err = global.GVA_DB.Find(&sysDictionarys).Error
	return sysDictionarys, err
}

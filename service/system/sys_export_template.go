package system

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
	systemReq "gin-vue-admin/model/system/request"
	"gin-vue-admin/utils"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"mime/multipart"
	"net/url"
	"strconv"
	"strings"
	"time"
)

type SysExportTemplateService struct{}

// CreateSysExportTemplate
//
//	@function:		CreateSysExportTemplate
//	@description:	创建导出模板记录
//	@param:			sysExportTemplate *system.SysExportTemplate
//	@return:		err error
func (sysExportTemplateService *SysExportTemplateService) CreateSysExportTemplate(sysExportTemplate *system.SysExportTemplate) (err error) {
	err = global.GVA_DB.Create(sysExportTemplate).Error
	return err
}

// DeleteSysExportTemplate
//
//	@function:		DeleteSysExportTemplate
//	@description:	删除导出模板记录
//	@param:			sysExportTemplate system.SysExportTemplate
//	@return:		err error
func (sysExportTemplateService *SysExportTemplateService) DeleteSysExportTemplate(sysExportTemplate system.SysExportTemplate) (err error) {
	err = global.GVA_DB.Delete(&sysExportTemplate).Error
	return err
}

// DeleteSysExportTemplateByIds
//
//	@function:		DeleteSysExportTemplateByIds
//	@description:	批量删除导出模板记录
//	@param:			ids request.IdsReq
//	@return:		err error
func (sysExportTemplateService *SysExportTemplateService) DeleteSysExportTemplateByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysExportTemplate{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSysExportTemplate
//
//	@function:		UpdateSysExportTemplate
//	@description:	更新导出模板记录
//	@param:			sysExportTemplate system.SysExportTemplate
//	@return:		err error
func (sysExportTemplateService *SysExportTemplateService) UpdateSysExportTemplate(sysExportTemplate system.SysExportTemplate) (err error) {
	return global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		conditions := sysExportTemplate.Conditions
		e := tx.Delete(&[]system.Condition{}, "template_id = ?", sysExportTemplate.TemplateID).Error
		if e != nil {
			return e
		}
		sysExportTemplate.Conditions = nil
		joins := sysExportTemplate.JoinTemplate
		e = tx.Delete(&[]system.JoinTemplate{}, "template_id = ?", sysExportTemplate.TemplateID).Error
		if e != nil {
			return e
		}
		sysExportTemplate.JoinTemplate = nil
		e = tx.Updates(&sysExportTemplate).Error
		if e != nil {
			return e
		}
		if len(conditions) > 0 {
			for i := range conditions {
				conditions[i].ID = 0
			}
			e = tx.Create(&conditions).Error
		}
		if len(joins) > 0 {
			for i := range joins {
				joins[i].ID = 0
			}
			e = tx.Create(&joins).Error
		}
		return e
	})
}

// GetSysExportTemplate
//
//	@function:		GetSysExportTemplate
//	@description:	根据id获取导出模板记录
//	@param:			id uint
//	@return:		sysExportTemplate system.SysExportTemplate, err error
func (sysExportTemplateService *SysExportTemplateService) GetSysExportTemplate(id uint) (sysExportTemplate system.SysExportTemplate, err error) {
	err = global.GVA_DB.Where("id = ?", id).Preload("JoinTemplate").Preload("Conditions").First(&sysExportTemplate).Error
	return
}

// GetSysTemplateInfoList
//
//	@function:		GetSysTemplateInfoList
//	@description:	分页获取导出模板记录
//	@param:			info systemReq.SysExportTemplateSearch
//	@return:		list []system.SysExportTemplate, total int64, err error
func (sysExportTemplateService *SysExportTemplateService) GetSysTemplateInfoList(info systemReq.SysExportTemplateSearch) (list []system.SysExportTemplate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysExportTemplate{})
	var sysExportTemplates []system.SysExportTemplate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.TableName != "" {
		db = db.Where("table_name LIKE ?", "%"+info.TableName+"%")
	}
	if info.TemplateID != "" {
		db = db.Where("template_id = ?", info.TemplateID)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}
	err = db.Find(&sysExportTemplates).Error
	return sysExportTemplates, total, err
}

// ExportExcel
//
//	@function:		ExportExcel
//	@description:	导出Excel
//	@param:			templateID string, value url.Values
//	@return:		file *bytes.Buffer, name string, err error
func (sysExportTemplateService *SysExportTemplateService) ExportExcel(templateID string, values url.Values) (file *bytes.Buffer, name string, err error) {
	var template system.SysExportTemplate
	err = global.GVA_DB.Preload("Conditions").Preload("JoinTemplate").First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return nil, "", err
	}
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 创建一个sheet
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var templateInfoMap = make(map[string]string)
	columns, err := utils.GetJSONKeys(template.TemplateInfo)
	if err != nil {
		return nil, "", err
	}
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return nil, "", err
	}
	var tableTitle []string
	for _, key := range columns {
		tableTitle = append(tableTitle, templateInfoMap[key])
	}
	selects := strings.Join(columns, ", ")
	var tableMap []map[string]interface{}
	db := global.GVA_DB
	if template.DBName != "" {
		db = global.MustGetGlobalDBByDBName(template.DBName)
	}
	if len(template.JoinTemplate) > 0 {
		for _, joinTemplate := range template.JoinTemplate {
			db = db.Joins(joinTemplate.JOINS + "`" + joinTemplate.Table + "`" + "ON" + joinTemplate.ON)
		}
	}
	db = db.Select(selects).Table(template.TableName)
	if len(template.Conditions) > 0 {
		for _, condition := range template.Conditions {
			sql := fmt.Sprintf("%s %s ?", condition.Column, condition.Operator)
			value := values.Get(condition.Form)
			if value != "" {
				if condition.Operator == "LIKE" {
					value = "%" + value + "%"
				}
				db = db.Where(sql, value)
			}
		}
	}
	// 通过参数传入limit
	limit := values.Get("limit")
	if limit != "" {
		l, e := strconv.Atoi(limit)
		if e == nil {
			db = db.Limit(l)
		}
	}
	// 模板默认limit
	if limit == "" && template.Limit != 0 {
		db = db.Limit(template.Limit)
	}
	// 通过参数传入offset
	offset := values.Get("offset")
	if offset != "" {
		o, e := strconv.Atoi(offset)
		if e == nil {
			db = db.Offset(o)
		}
	}
	// 通过参数传入order
	order := values.Get("order")
	if order != "" {
		db = db.Order(order)
	}
	// 模板默认order
	if order == "" && template.Order != "" {
		db = db.Order(template.Order)
	}
	err = db.Debug().Find(&tableMap).Error
	if err != nil {
		return nil, "", err
	}
	var rows [][]string
	rows = append(rows, tableTitle)
	for _, table := range tableMap {
		var row []string
		for _, column := range columns {
			if len(template.JoinTemplate) > 0 {
				columnArr := strings.Split(column, ".")
				if len(columnArr) > 1 {
					column = strings.Split(column, ".")[1]
				}
			}
			row = append(row, fmt.Sprintf("%v", table[column]))
		}
		rows = append(rows, row)
	}
	for i, row := range rows {
		for j, colCell := range row {
			err := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", getColumnName(j+1), i+1), colCell)
			if err != nil {
				return nil, "", err
			}
		}
	}
	f.SetActiveSheet(index)
	file, err = f.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}
	return file, template.Name, nil
}

// ExportTemplate
//
//	@function:		ExportTemplate
//	@description:	导出Excel模板
//	@param:			templateID string
//	@return:		file *bytes.Buffer, name string, err error
func (sysExportTemplateService *SysExportTemplateService) ExportTemplate(templateID string) (file *bytes.Buffer, name string, err error) {
	var template system.SysExportTemplate
	err = global.GVA_DB.First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return nil, "", err
	}
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 创建新的sheet
	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	var templateInfoMap = make(map[string]string)
	columns, err := utils.GetJSONKeys(template.TemplateInfo)
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return nil, "", err
	}
	var tableTitle []string
	for _, key := range columns {
		tableTitle = append(tableTitle, templateInfoMap[key])
	}
	for i := range tableTitle {
		fErr := f.SetCellValue("Sheet1", fmt.Sprintf("%s%d", getColumnName(i+1), 1), tableTitle[i])
		if fErr != nil {
			return nil, "", fErr
		}
	}
	f.SetActiveSheet(index)
	file, err = f.WriteToBuffer()
	if err != nil {
		return nil, "", err
	}
	return file, template.Name, nil
}

// ImportExcel
//
//	@function:		ImportExcel
//	@description:	导入Excel
//	@param:			templateID string, file *multipart.FileHeader
//	@return:		err error
func (sysExportTemplateService *SysExportTemplateService) ImportExcel(templateID string, file *multipart.FileHeader) (err error) {
	var template system.SysExportTemplate
	err = global.GVA_DB.First(&template, "template_id = ?", templateID).Error
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	f, err := excelize.OpenReader(src)
	if err != nil {
		return err
	}
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}
	var templateInfoMap = make(map[string]string)
	err = json.Unmarshal([]byte(template.TemplateInfo), &templateInfoMap)
	if err != nil {
		return err
	}
	var titleKeyMap = make(map[string]string)
	for key, title := range templateInfoMap {
		titleKeyMap[title] = key
	}
	db := global.GVA_DB
	if template.DBName != "" {
		db = global.MustGetGlobalDBByDBName(template.DBName)
	}
	return db.Transaction(func(tx *gorm.DB) error {
		excelTitle := rows[0]
		values := rows[1:]
		items := make([]map[string]interface{}, 0, len(values))
		for _, row := range values {
			var item = make(map[string]interface{})
			for i, value := range row {
				key := titleKeyMap[excelTitle[i]]
				item[key] = value
			}
			needCreated := tx.Migrator().HasColumn(template.TableName, "created_at")
			needUpdated := tx.Migrator().HasColumn(template.TableName, "updated_at")

			if item["created_at"] == nil && needCreated {
				item["created_at"] = time.Now()
			}
			if item["updated_at"] == nil && needUpdated {
				item["updated_at"] = time.Now()
			}
			items = append(items, item)
		}
		cErr := tx.Table(template.TableName).CreateInBatches(&items, 1000).Error
		return cErr
	})
}

func getColumnName(n int) string {
	columnName := ""
	for n > 0 {
		n--
		columnName = string(rune('A'+n%26)) + columnName
		n /= 26
	}
	return columnName
}

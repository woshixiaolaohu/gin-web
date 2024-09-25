package example

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/example"
	"gin-vue-admin/utils/upload"
	"mime/multipart"
	"strings"
)

// Upload
//
//	@author:		jelly
//	@function:		Upload
//	@description:	创建文件上传记录
//	@param:			file model.ExaFileUploadAndDownload
//	@return:		error
func (f *FileUploadAndDownloadService) Upload(file example.ExaFileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}

// FindFile
//
//	@author:		jelly
//	@function:		FindFile
//	@description:	查询文件记录
//	@param:			id uint
//	@return:		model.ExaFileUploadAndDownload, error
func (f *FileUploadAndDownloadService) FindFile(id uint) (example.ExaFileUploadAndDownload, error) {
	var file example.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

// DeleteFile
//
//	@author:		jelly
//	@function:		DeleteFile
//	@description:	删除文件记录
//	@param:			file model.ExaFileUploadAndDownload
//	@return:		err error
func (f *FileUploadAndDownloadService) DeleteFile(file example.ExaFileUploadAndDownload) (err error) {
	var fileFormDb example.ExaFileUploadAndDownload
	fileFormDb, err = f.FindFile(file.ID)
	if err != nil {
		return err
	}
	oss := upload.NewOss()
	if err = oss.DeleteFile(fileFormDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.GVA_DB.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

// EditFileName 编辑文件名或者备注
func (f *FileUploadAndDownloadService) EditFileName(file example.ExaFileUploadAndDownload) (err error) {
	var fileFormDb example.ExaFileUploadAndDownload
	return global.GVA_DB.Where("id = ?", file.ID).First(&fileFormDb).Update("name", file.Name).Error
}

// GetFileRecordInfoList
//
//	@author:		jelly
//	@function:		GetFileRecordInfoList
//	@description:	分页获取数据
//	@param:			info request.PageInfo
//	@return:		list interface{}, total int64, err error
func (f *FileUploadAndDownloadService) GetFileRecordInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.KeyWord
	db := global.GVA_DB.Model(&example.ExaFileUploadAndDownload{})
	var fileLists []example.ExaFileUploadAndDownload
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?" + "%" + keyword + "%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}

// UploadFile
//
//	@author:		jelly
//	@function:		UploadFile
//	@description:	根据配置文件判断是文件上传到本地还是云
//	@param:			header *multipart.FileHeader, noSave string
//	@return:		file model.ExaFileUploadAndDownload, err error
func (e *FileUploadAndDownloadService) UploadFile(header *multipart.FileHeader, noSave string) (file example.ExaFileUploadAndDownload, err error) {
	oss := upload.NewOss()
	filePath, key, uploadErr := oss.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}
	s := strings.Split(header.Filename, ".")
	f := example.ExaFileUploadAndDownload{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	if noSave == "0" {
		return f, e.Upload(f)
	}
	return f, nil
}

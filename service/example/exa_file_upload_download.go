package example

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/example"
	"gin-vue-admin/utils/upload"
)

// Upload
// @author: jelly
// @function: Upload
// @description: 创建文件上传记录
// @param: file model.ExaFileUploadAndDownload
// @return: error
func (f *FileUploadAndDownloadService) Upload(file example.ExaFileUploadAndDownload) error {
	return global.GVA_DB.Create(&file).Error
}

// FindFile
// @author: jelly
// @function: FindFile
// @description: 查询文件记录
// @param: id uint
// @return: model.ExaFileUploadAndDownload, error
func (f *FileUploadAndDownloadService) FindFile(id uint) (example.ExaFileUploadAndDownload, error) {
	var file example.ExaFileUploadAndDownload
	err := global.GVA_DB.Where("id = ?", id).First(&file).Error
	return file, err
}

// DeleteFile
// @author: jelly
// @function: DeleteFile
// @description: 删除文件记录
// @param: file model.ExaFileUploadAndDownload
// @return: err error
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

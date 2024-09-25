package upload

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/utils"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Local struct {
}

// UploadFile
//
//	@author:		jelly
//	@object:		*Local
//	@function:		UploadFile
//	@description:	上传文件
//	@param:			file *multipart.FileHeader
//	@return:		string, string, error
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	// 读取文件后缀
	ext := filepath.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新的文件名
	filename := name + "_" + time.Now().Format("220060102150405") + ext
	// 尝试创建此路径
	mkdirErr := os.MkdirAll(global.GVA_CONFIG.Local.StorePath, os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("function os.MkdirAll() failed", zap.Any("err", mkdirErr.Error()))
		return "", "", errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
	}
	// 拼接路径和文件名
	p := global.GVA_CONFIG.Local.StorePath + "/" + filename
	filePath := global.GVA_CONFIG.Local.Path + "/" + filename
	// 读取文件
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	// 创建文件 defer 关闭
	defer f.Close()
	out, createError := os.Create(p)
	if createError != nil {
		global.GVA_LOG.Error("function os.Create() failed", zap.Any("err", createError.Error()))
		return "", "", errors.New("function os.Create() failed, err:" + createError.Error())
	}
	// 创建文件 defer 关闭
	defer out.Close()
	// 传输（拷贝）文件
	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() failed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() failed, err:" + copyErr.Error())
	}
	return filePath, filename, nil
}

// DeleteFile
//
//	@author:		jelly
//	@object:		*Local
//	@function:		DeleteFile
//	@description:	删除文件
//	@param:			key string
//	@return:		error
func (*Local) DeleteFile(key string) error {
	p := global.GVA_CONFIG.Local.StorePath + "/" + key
	if strings.Contains(p, global.GVA_CONFIG.Local.StorePath) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败,err:" + err.Error())
		}
	}
	return nil
}

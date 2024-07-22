package upload

import (
	"errors"
	"gin-vue-admin/global"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
	"mime/multipart"
	"time"
)

type AliyunOSS struct {
}

func (*AliyunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() failed, err:" + err.Error())
	}
	// 读取本地文件
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	// 创建文件 defer 关闭
	defer f.Close()
	// 上传阿里云路径 文件名格式 保持唯一性
	yunFileTempPath := global.GVA_CONFIG.AliyunOSS.BasePath + "/" + "upload" + "/" + time.Now().Format("2006-01-02") + "/" + file.Filename
	//上传文件流
	err = bucket.PutObject(yunFileTempPath, f)
	if err != nil {
		global.GVA_LOG.Error("function formUploader.Put() failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() failed, err:" + err.Error())
	}
	return global.GVA_CONFIG.AliyunOSS.BucketUrl + "/" + yunFileTempPath, yunFileTempPath, nil
}

func (*AliyunOSS) DeleteFile(key string) error {
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() failed", zap.Any("err", err.Error()))
		return errors.New("function AliyunOSS.NewBucket() failed, err:" + err.Error())
	}
	// 删除单个文件 objectName 表示删除 OSS 文件时需要指定包含文件后缀在内的完整路径 例如 abc/efg/123.jpg
	// 如需删除文件夹 将 objectName 设置为对应的文件夹名称 如果文件夹非空 则需要将文件夹下的所有 object 删除后才能删除该文件夹
	err = bucket.DeleteObject(key)
	if err != nil {
		global.GVA_LOG.Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}

func NewBucket() (*oss.Bucket, error) {
	// 创建 OSSClient 实例
	client, err := oss.New(global.GVA_CONFIG.AliyunOSS.EndPoint, global.GVA_CONFIG.AliyunOSS.AccessKeyID, global.GVA_CONFIG.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	// 获取存储空间
	bucket, err := client.Bucket(global.GVA_CONFIG.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}
	return bucket, err
}

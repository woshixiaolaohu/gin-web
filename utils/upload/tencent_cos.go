package upload

import (
	"context"
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"github.com/tencentyun/cos-go-sdk-v5"
	"go.uber.org/zap"
	"mime/multipart"
	"net/http"
	"net/url"
	"time"
)

type TencentCOS struct{}

// UploadFile 上传文件
func (t *TencentCOS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	client := NewClient()
	f, openError := file.Open()
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)

	_, err := client.Object.Put(context.Background(), global.GVA_CONFIG.TencentCOS.PathPrefix+"/"+fileKey, f, nil)
	if err != nil {
		panic(err)
	}
	return global.GVA_CONFIG.TencentCOS.BaseURL + "/" + global.GVA_CONFIG.TencentCOS.PathPrefix + "/" + fileKey, fileKey, nil
}
func NewClient() *cos.Client {
	urlStr, _ := url.Parse("https://" + global.GVA_CONFIG.TencentCOS.Bucket + ".cos." + global.GVA_CONFIG.TencentCOS.Region + ".mycloud.com")
	baseUrl := &cos.BaseURL{BucketURL: urlStr}
	client := cos.NewClient(baseUrl, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  global.GVA_CONFIG.TencentCOS.SecretID,
			SecretKey: global.GVA_CONFIG.TencentCOS.SecretKey,
		},
	})
	return client
}

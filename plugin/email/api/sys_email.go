package api

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"
	emailresponse "gin-vue-admin/plugin/email/model/response"
	"gin-vue-admin/plugin/email/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EmailApi struct{}

// EmailTest
//
//	@Tags		System
//	@Summary	发送测试邮件
//	@Security	ApiKeyAuth
//	@Produce	application/json
//	@Success	200	{string}	string	"{"success":true,"data":{},"msg":"发送成功"}"
//	@Router		/email/emailTest [post]
func (s *EmailApi) EmailTest(c *gin.Context) {
	err := service.ServiceGroupApp.EmailTest()
	if err != nil {
		global.GVA_LOG.Error("发送失败", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

func (s *EmailApi) SendEmail(c *gin.Context) {
	var email emailresponse.Email
	err := c.ShouldBindJSON(&email)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = service.ServiceGroupApp.SendEmail(email.To, email.Subject, email.Body)
	if err != nil {
		global.GVA_LOG.Error("发送失败", zap.Error(err))
		response.FailWithMessage("发送失败", c)
		return
	}
	response.OkWithMessage("发送成功", c)
}

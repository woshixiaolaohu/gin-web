package system

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/response"
	"gin-vue-admin/model/system"
	"gin-vue-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type JwtApi struct{}

// JsonInBlacklist
//
//	@Tags		Jwt
//	@Summary	jwt加入黑名单
//	@Security	ApiKeyAuth
//	@accept		application/json
//	@Produce	application/json
//	@Success	200	{object}	response.Response{msg=string}	"jwt加入黑名单"
//	@Router		/jwt/jsonInBlacklist [post]
func (s *JwtApi) JsonInBlacklist(c *gin.Context) {
	token := utils.GetToken(c)
	jwt := system.JwtBlacklist{Jwt: token}
	err := jwtService.JsonInBlackList(jwt)
	if err != nil {
		global.GVA_LOG.Error("jwt作废失败", zap.Error(err))
		response.FailWithMessage("jwt作废失败", c)
		return
	}
	utils.ClearToken(c)
	response.OkWithMessage("jwt作废成功", c)
}

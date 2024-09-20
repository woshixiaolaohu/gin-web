package system

import "github.com/gin-gonic/gin"
import v1 "gin-vue-admin/api"

type AuthorityBtnRouter struct{}

func (s *AuthorityBtnRouter) InitAuthorityBtnRouterRouter(Router *gin.RouterGroup) {
	authorityRouterWithoutRecord := Router.Group("authorityBtn")
	authorityBtnApi := v1.ApiGroupApp.SystemApiGroup.AuthorityBtnApi
	{
		authorityRouterWithoutRecord.POST("getAuthorityBtn", authorityBtnApi.GetAuthorityBtn)
		authorityRouterWithoutRecord.POST("setAuthorityBtn", authorityBtnApi.SetAuthorityBtn)
		authorityRouterWithoutRecord.POST("canRemoveAuthorityBtn", authorityBtnApi.CanRemoveAuthorityBtn)
	}
}

package example

import (
	v1 "gin-vue-admin/api"
	"gin-vue-admin/middleware"
	"github.com/gin-gonic/gin"
)

type CustomerRouter struct {
}

func (c *CustomerRouter) InitCustomRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.OperationRecord())
	customerRouterWithoutRecord := Router.Group("customer")
	exaCustomerApi := v1.ApiGroupApp.ExampleApiGroup.CustomerApi
	{
		customerRouter.POST("createExaCustomer", exaCustomerApi.CreateExaCustomer)   // 创建客户
		customerRouter.PUT("updateExaCustomer", exaCustomerApi.UpdateExaCustomer)    // 更新客户
		customerRouter.DELETE("deleteExaCustomer", exaCustomerApi.DeleteExaCustomer) // 删除客户
	}
	{
		customerRouterWithoutRecord.GET("getExaCustomer", exaCustomerApi.GetExaCustomer)         // 获取单一客户信息
		customerRouterWithoutRecord.GET("getExaCustomerList", exaCustomerApi.GetExaCustomerList) // 获取客户列表
	}

}

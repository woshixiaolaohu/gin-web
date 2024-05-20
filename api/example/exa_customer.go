package example

import (
	"gin-vue-admin/model/common/response"
	"gin-vue-admin/model/example"
	"github.com/gin-gonic/gin"
)

type CustomerApi struct{}

// CreateExaCustomer 创建客户
func (e *CustomerApi) CreateExaCustomer(c *gin.Context) {
	var customer example.ExaCustomer
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

}

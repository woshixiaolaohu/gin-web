package example

import "github.com/gin-gonic/gin"

type CustomerRouter struct {
}

func (c *CustomerRouter) InitCustomRouter(Router *gin.RouterGroup) {
	customerRouter := Router.Group("customer").Use(middleware.O)
}

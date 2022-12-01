package controller

import (
	"github.com/gin-gonic/gin"
	"jwt-demo/internal/controller/account"
	"jwt-demo/internal/controller/test"
)

func Routes(rg *gin.RouterGroup) {
	accountRoutes := rg.Group("/account")
	{
		account.Routes(accountRoutes)
	}

	testRoutes := rg.Group("/test")
	{
		test.RoutesTest(testRoutes)
	}
}

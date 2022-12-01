package test

import (
	"github.com/gin-gonic/gin"
	"jwt-demo/internal/middleware"
	"net/http"
)

func RoutesTest(r *gin.RouterGroup) {
	r.Use(middleware.AuthorizeJWT())
	{
		r.GET("/", func(c *gin.Context) {
			userName, _ := c.Get(middleware.ContextKeySub.String())
			c.JSON(http.StatusOK, gin.H{
				"message":   "ok",
				"user_name": userName,
			})
		})
	}
}

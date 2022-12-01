package account

import (
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.RouterGroup) {
	r.POST("/login", func(c *gin.Context) {
		Login(c)
	})
}

package account

import (
	"github.com/gin-gonic/gin"
	"jwt-demo/internal/httpbody/request"
	"jwt-demo/internal/service/account"
	"net/http"
)

func Login(c *gin.Context) {
	var loginRequest request.Login
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	loginResponse, err := account.Login(c, loginRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, loginResponse)
}

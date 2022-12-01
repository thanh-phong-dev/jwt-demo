package account

import (
	"github.com/gin-gonic/gin"
	"jwt-demo/internal/httpbody/request"
	"jwt-demo/internal/httpbody/response"
	"jwt-demo/internal/middleware"
	"jwt-demo/internal/utils"
)

// hash password for value "1234"
const hashPassword = "$2a$10$8584ArWGmm2LvO8BTPvePOb3ea2.cUXxrM.Q7NREiXOJKZVyWhOcS"

func Login(c *gin.Context, req request.Login) (*response.Login, error) {
	if err := utils.CheckPasswordHash(req.PassWord, hashPassword); err != nil {
		return nil, err
	}

	token, err := middleware.GenerateToken(req.UserName)
	if err != nil {
		return nil, err
	}

	return &response.Login{Token: token}, nil
}

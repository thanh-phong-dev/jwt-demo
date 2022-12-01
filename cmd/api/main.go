package main

import (
	"github.com/gin-gonic/gin"
	"jwt-demo/internal/controller"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/api")
	{
		controller.Routes(v1)
	}
	
	router.Run(":8080")
}

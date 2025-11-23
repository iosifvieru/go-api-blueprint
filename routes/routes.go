package routes

import (
	"go-api-blueprint/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.GET("/health", controllers.GetServiceHealth)
	}
}

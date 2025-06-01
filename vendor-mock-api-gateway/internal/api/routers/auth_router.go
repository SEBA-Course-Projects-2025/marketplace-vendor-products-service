package routers

import (
	"github.com/gin-gonic/gin"
	"vendor-mock-api-gateway/vendor-mock-api-gateway/internal/api/handlers"
)

func SetUpAuthRouter(rg *gin.RouterGroup) {
	auth := rg.Group("auth")
	{
		auth.POST("/registration", handlers.PostRegistrationHandler)
		auth.POST("/login", handlers.PostLoginHandler)
	}
}

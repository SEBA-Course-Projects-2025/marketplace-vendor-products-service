package routers

import (
	"customer-mock-api-gateway/customer-mock-api-gateway/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpAuthRouter(rg *gin.RouterGroup) {

	auth := rg.Group("auth")
	{
		auth.POST("/registration", handlers.PostRegistrationHandler)
		auth.POST("/login", handlers.PostLoginHandler)
	}

}

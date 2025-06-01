package routers

import (
	"github.com/gin-gonic/gin"
	"mock-api-gateway/mock-api-gateways/customer-mock-api-gateway/internal/api/handlers"
)

func SetUpProfileRouter(rg *gin.RouterGroup) {

	profiles := rg.Group("customers")
	{
		profiles.GET("/:customerId", handlers.GetProfileHandler)
		profiles.PUT("/:customerId", handlers.PutProfileHandler)
	}

}

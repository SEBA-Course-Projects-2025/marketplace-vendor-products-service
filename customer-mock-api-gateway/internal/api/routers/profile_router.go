package routers

import (
	"customer-mock-api-gateway/customer-mock-api-gateway/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpProfileRouter(rg *gin.RouterGroup) {

	profiles := rg.Group("customers")
	{
		profiles.GET("/:customerId", handlers.GetProfileHandler)
		profiles.PUT("/:customerId", handlers.PutProfileHandler)
	}

}

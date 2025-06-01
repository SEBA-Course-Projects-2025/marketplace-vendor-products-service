package routers

import (
	"github.com/gin-gonic/gin"
	"mock-api-gateway/mock-api-gateways/vendor-mock-api-gateway/internal/api/handlers"
)

func SetUpProfileRouter(rg *gin.RouterGroup) {
	profiles := rg.Group("vendors")
	{
		profiles.GET("/:vendorId", handlers.GetProfileHandler)
		profiles.PUT("/:vendorId", handlers.PutProfileHandler)
	}
}

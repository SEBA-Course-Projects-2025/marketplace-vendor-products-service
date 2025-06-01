package routers

import (
	"github.com/gin-gonic/gin"
	"mock-api-gateway/mock-api-gateways/customer-mock-api-gateway/internal/api/middlewares"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	rootGroup := r.Group("/")

	SetUpAuthRouter(rootGroup)

	{
		rootGroup.Use(middlewares.AuthenticationMiddleware())

		SetUpProfileRouter(rootGroup)
		SetUpCartRouter(rootGroup)
		SetUpLikedProductsRouter(rootGroup)

	}

	return r

}

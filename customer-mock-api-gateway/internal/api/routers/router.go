package routers

import (
	"customer-mock-api-gateway/customer-mock-api-gateway/internal/api/middlewares"
	"github.com/gin-gonic/gin"
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

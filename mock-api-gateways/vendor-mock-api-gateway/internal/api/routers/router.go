package routers

import (
	"github.com/gin-gonic/gin"
	"mock-api-gateway/mock-api-gateways/vendor-mock-api-gateway/internal/api/middlewares"
)

func SetUpRouter() *gin.Engine {

	r := gin.Default()

	rootGroup := r.Group("/")

	SetUpAuthRouter(rootGroup)

	{
		rootGroup.Use(middlewares.AuthenticationMiddleware())

		SetUpProfileRouter(rootGroup)
		SetUpProductsRouter(rootGroup)
		SetUpProductsStockRouter(rootGroup)
		SetUpReviewsRouter(rootGroup)
		SetUpOrdersRouter(rootGroup)

	}

	return r

}

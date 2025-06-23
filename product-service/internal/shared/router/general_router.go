package router

import (
	_ "dev-vendor/docs"
	productInterfaces "dev-vendor/product-service/internal/products/interfaces"
	productsHandlers "dev-vendor/product-service/internal/products/interfaces/handlers"
	"dev-vendor/product-service/internal/shared/middlewares"
	stocksInterfaces "dev-vendor/product-service/internal/stocks/interfaces"
	stockHandlers "dev-vendor/product-service/internal/stocks/interfaces/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetUpRouter(productHandler *productsHandlers.ProductHandler, stockHandler *stockHandlers.StockHandler) *gin.Engine {

	r := gin.New()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api", middlewares.AuthMiddleware())

	productInterfaces.SetUpProductsRouter(api, productHandler)
	stocksInterfaces.SetUpStocksRouter(api, stockHandler)

	return r
}

package router

import (
	productInterfaces "dev-vendor/product-service/internal/products/interfaces"
	productsHandlers "dev-vendor/product-service/internal/products/interfaces/handlers"
	"dev-vendor/product-service/internal/shared/middlewares"
	stocksInterfaces "dev-vendor/product-service/internal/stocks/interfaces"
	stockHandlers "dev-vendor/product-service/internal/stocks/interfaces/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpRouter(productHandler *productsHandlers.ProductHandler, stockHandler *stockHandlers.StockHandler) *gin.Engine {

	r := gin.New()

	api := r.Group("/api", middlewares.VendorIdMiddleware())

	productInterfaces.SetUpProductsRouter(api, productHandler)
	stocksInterfaces.SetUpStocksRouter(api, stockHandler)

	return r
}

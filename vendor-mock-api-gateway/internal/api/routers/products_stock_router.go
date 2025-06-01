package routers

import (
	"github.com/gin-gonic/gin"
	"vendor-mock-api-gateway/vendor-mock-api-gateway/internal/api/handlers"
)

func SetUpProductsStockRouter(rg *gin.RouterGroup) {
	productsStock := rg.Group("products/stock")
	{
		productsStock.GET("/", handlers.GetStockProductsHandler)

		productsStock.GET("/:stockId", handlers.GetOneStockProductHandler)
		productsStock.PUT("/:stockId", handlers.PutStockProductHandler)
	}
}

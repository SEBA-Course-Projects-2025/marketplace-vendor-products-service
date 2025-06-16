package interfaces

import (
	"dev-vendor/product-service/internal/stocks/interfaces/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpStocksRouter(rg *gin.RouterGroup, h *handlers.StockHandler) {
	stocks := rg.Group("stocks")
	{
		stocks.GET("/", h.GetAllStocksHandler)
		stocks.POST("/", h.PostStockHandler)
		stocks.PATCH("/:stockId/products", h.PatchStockProductsHandler)
		stocks.DELETE("/", h.DeleteManyStocksHandler)
		stocks.DELETE("/:stockId/products", h.DeleteManyStockProductsHandler)

		stocks.GET("/:stockId", h.GetStockHandler)
		stocks.PUT("/:stockId", h.PutStockHandler)
		stocks.PUT("/:stockId/products/:productId", h.PutStockProductHandler)
		stocks.PATCH("/:stockId", h.PatchStockByIdHandler)
		stocks.PATCH("/:stockId/products/:productId", h.PatchStockProductByIdHandler)
		stocks.DELETE("/:stockId", h.DeleteStockByIdHandler)
		stocks.DELETE("/:stockId/products/:productId", h.DeleteStockProductByIdHandler)
	}
}

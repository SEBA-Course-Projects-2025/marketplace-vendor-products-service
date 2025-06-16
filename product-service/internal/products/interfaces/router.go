package interfaces

import (
	"dev-vendor/product-service/internal/products/interfaces/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpProductsRouter(rg *gin.RouterGroup, h *handlers.ProductHandler) {
	products := rg.Group("products")
	{
		products.GET("/", h.GetAllProductsHandler)
		products.POST("/", h.PostProductHandler)
		products.DELETE("/", h.DeleteManyProductsHandler)

		products.GET("/:productId", h.GetProductHandler)
		products.PUT("/:productId", h.PutProductHandler)
		products.PATCH("/:productId", h.PatchProductHandler)
		products.DELETE("/:productId", h.DeleteProductHandler)
	}
}

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

		products.GET("/id/:productId", h.GetProductHandler)
		products.PUT("/id/:productId", h.PutProductHandler)
		products.PATCH("/id/:productId", h.PatchProductHandler)
		products.DELETE("/id/:productId", h.DeleteProductHandler)

		products.GET("/slug/:slug", h.GetProductBySlugHandler)
		products.PUT("/slug/:slug", h.PutProductBySlugHandler)
		products.PATCH("/slug/:slug", h.PatchProductBySlugHandler)
		products.DELETE("/slug/:slug", h.DeleteProductBySlugHandler)

	}
}

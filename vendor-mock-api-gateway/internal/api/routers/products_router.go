package routers

import (
	"github.com/gin-gonic/gin"
	"vendor-mock-api-gateway/vendor-mock-api-gateway/internal/api/handlers"
)

func SetUpProductsRouter(rg *gin.RouterGroup) {
	products := rg.Group("products")
	{
		products.GET("/", handlers.GetProductsHandler)
		products.POST("/", handlers.PostProductsHandler)
		products.PATCH("/", handlers.PatchProductsHandler)
		products.DELETE("/", handlers.DeleteProductsHandler)

		products.GET("/:productId", handlers.GetOneProductHandler)
		products.PUT("/:productId", handlers.PutProductHandler)
		products.PATCH("/:productId", handlers.PatchOneProductHandler)
		products.DELETE("/:productId", handlers.DeleteOneProductHandler)
	}
}

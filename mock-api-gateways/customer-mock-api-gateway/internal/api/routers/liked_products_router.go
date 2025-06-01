package routers

import (
	"github.com/gin-gonic/gin"
	"mock-api-gateway/mock-api-gateways/customer-mock-api-gateway/internal/api/handlers"
)

func SetUpLikedProductsRouter(rg *gin.RouterGroup) {

	likedProducts := rg.Group("liked-products")
	{
		likedProducts.GET("/", handlers.GetLikedProductsHandler)
		likedProducts.DELETE("/", handlers.DeleteLikedProductsHandler)

		likedProducts.PUT("/:likedProductId", handlers.PutLikedProductsHandler)
		likedProducts.DELETE("/:likedProductId", handlers.DeleteOneLikedProductHandler)
	}

}

package routers

import (
	"customer-mock-api-gateway/customer-mock-api-gateway/internal/api/handlers"
	"github.com/gin-gonic/gin"
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

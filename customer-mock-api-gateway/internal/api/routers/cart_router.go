package routers

import (
	"customer-mock-api-gateway/customer-mock-api-gateway/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetUpCartRouter(rg *gin.RouterGroup) {

	cart := rg.Group("cart")
	{
		cart.GET("/", handlers.GetCartHandler)
		cart.POST("/items", handlers.PostCartHandler)

		cart.PUT("/items/:productId", handlers.PutCartQuantityHandler)
		cart.DELETE("/items/:productId", handlers.DeleteCartProductHandler)
	}

	rg.POST("/checkout", handlers.PostCheckoutHandler)
}

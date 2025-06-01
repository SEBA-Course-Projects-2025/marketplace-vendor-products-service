package routers

import (
	"github.com/gin-gonic/gin"
	"vendor-mock-api-gateway/vendor-mock-api-gateway/internal/api/handlers"
)

func SetUpOrdersRouter(rg *gin.RouterGroup) {
	orders := rg.Group("orders")
	{
		orders.GET("/", handlers.GetOrdersHandler)

		orders.GET("/:orderId", handlers.GetOneOrderHandler)
		orders.PUT("/:orderId", handlers.PutOrderHandler)
	}
}

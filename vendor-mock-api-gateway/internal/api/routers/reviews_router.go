package routers

import (
	"github.com/gin-gonic/gin"
	"vendor-mock-api-gateway/vendor-mock-api-gateway/internal/api/handlers"
)

func SetUpReviewsRouter(rg *gin.RouterGroup) {
	reviews := rg.Group("reviews")
	{
		reviews.GET("/", handlers.GetReviewsHandler)
		reviews.POST("/:reviewId/replies", handlers.PostReplyHandler)

		reviews.GET("/:reviewId", handlers.GetOneReviewHandler)
		reviews.PUT("/:reviewId/replies/:replyId", handlers.PutReplyHandler)

	}
}

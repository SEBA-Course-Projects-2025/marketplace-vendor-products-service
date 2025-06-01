package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"vendor-mock-api-gateway/vendor-mock-api-gateway/internal/models"
)

func GetOneReviewHandler(c *gin.Context) {

	reviewId := c.Param("reviewId")

	review := models.Review{
		Id:           "mockReviewId",
		ProductId:    "mockProductId",
		ReviewerId:   "mockReviewerId",
		ReviewerName: "mockReviewerName",
		Rating:       4.5,
		Comment:      "mockComment",
		Date:         time.Now(),
		Reply: []models.Reply{
			{
				Id:          "mockReplyId",
				ReplierId:   "mockReplierId",
				ReplierName: "mockReplierName",
				Comment:     "mockReplyComment",
				Date:        time.Now(),
			},
		},
	}

	if reviewId != review.Id {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})

}

func GetReviewsHandler(c *gin.Context) {

	reviews := []models.Review{
		{
			Id:           "mockReviewId",
			ProductId:    "mockProductId",
			ReviewerId:   "mockReviewerId",
			ReviewerName: "mockReviewerName",
			Rating:       4.5,
			Comment:      "mockComment",
			Date:         time.Now(),
			Reply: []models.Reply{
				{
					Id:          "mockReplyId",
					ReplierId:   "mockReplierId",
					ReplierName: "mockReplierName",
					Comment:     "mockReplyComment",
					Date:        time.Now(),
				},
			},
		},
		{
			Id:           "mockReviewId2",
			ProductId:    "mockProductId2",
			ReviewerId:   "mockReviewerId2",
			ReviewerName: "mockReviewerName2",
			Rating:       4,
			Comment:      "mockComment2",
			Date:         time.Now(),
			Reply: []models.Reply{
				{
					Id:          "mockReplyId2",
					ReplierId:   "mockReplierId2",
					ReplierName: "mockReplierName2",
					Comment:     "mockReplyComment2",
					Date:        time.Now(),
				},
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": reviews})

}

func PostReplyHandler(c *gin.Context) {

	reviewId := c.Param("reviewId")

	if reviewId != "mockReviewId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var replyComment models.ReplyComment

	if err := c.ShouldBindJSON(&replyComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment data"})
		return
	}

	reply := models.Reply{
		Id:          "mockReplyId",
		ReplierId:   "mockReplierId",
		ReplierName: "mockReplierName",
		Comment:     replyComment.Comment,
		Date:        time.Now(),
	}

	c.JSON(http.StatusOK, gin.H{"data": reply})

}

func PutReplyHandler(c *gin.Context) {

	reviewId := c.Param("reviewId")
	replyId := c.Param("replyId")

	if reviewId != "mockReviewId" || replyId != "mockReplyId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var replyComment models.ReplyComment

	if err := c.ShouldBindJSON(&replyComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

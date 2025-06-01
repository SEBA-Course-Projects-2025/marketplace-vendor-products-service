package handlers

import (
	"github.com/gin-gonic/gin"
	"mock-api-gateway/mock-api-gateways/customer-mock-api-gateway/internal/models"
	"net/http"
)

func GetLikedProductsHandler(c *gin.Context) {

	likedProducts := []models.LikedProduct{
		{
			ProductId:    "mockProductId",
			Name:         "mockProductName",
			Price:        100,
			Availability: "available",
			Image:        "https://example.com/product.png",
		},
		{
			ProductId:    "mockProductId2",
			Name:         "mockProductName2",
			Price:        200,
			Availability: "available",
			Image:        "https://example.com/product2.png",
		},
		{
			ProductId:    "mockProductId3",
			Name:         "mockProductName3",
			Price:        300,
			Availability: "available",
			Image:        "https://example.com/product3.png",
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": likedProducts})

}

func PutLikedProductsHandler(c *gin.Context) {

	likedProductId := c.Param("likedProductId")

	if likedProductId != "mockLikedProductId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var liked models.LikedCheck

	if err := c.ShouldBindJSON(&liked); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid liked product data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func DeleteOneLikedProductHandler(c *gin.Context) {

	likedProductId := c.Param("likedProductId")

	if likedProductId != "mockLikedProductId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deletedId": likedProductId})

}

func DeleteLikedProductsHandler(c *gin.Context) {

	var ids models.IdsToDelete

	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid liked product data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deletedIds": ids.Ids})

}

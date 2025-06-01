package handlers

import (
	"customer-mock-api-gateway/customer-mock-api-gateway/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProfileHandler(c *gin.Context) {

	id := c.Param("customerId")

	profile := models.Profile{
		Id:              "mockCustomerId",
		Email:           "mockCustomer@example.com",
		ShippingAddress: "789 Maple Drive, Toronto, ON M4B 1B3, Canada",
		Orders: []string{
			"mockOrder1",
			"mockOrder2",
		},
	}

	if profile.Id == id {
		c.JSON(http.StatusOK, gin.H{
			"data": profile,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{})

}

func PutProfileHandler(c *gin.Context) {

	id := c.Param("customerId")

	var profile models.Profile

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile data"})
		return
	}

	if id != "mockCustomerId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

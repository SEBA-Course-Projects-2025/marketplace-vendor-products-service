package handlers

import (
	"customer-mock-api-gateway/customer-mock-api-gateway/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PostRegistrationHandler(c *gin.Context) {

	var registration models.Registration

	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid registration data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token":      "mock-jwt-token",
		"customerId": "mock-customer-id",
	})

}

func PostLoginHandler(c *gin.Context) {

	var login models.Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login data"})
		return
	}

	validLogin := models.Login{
		Email:    "mockCustomer@example.com",
		Password: "password123",
	}

	if login.Email == validLogin.Email && login.Password == validLogin.Password {
		c.JSON(http.StatusOK, gin.H{
			"tokenSession":      "mock-jwt-token",
			"customerSessionId": "mock-customer-session-id",
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})

}

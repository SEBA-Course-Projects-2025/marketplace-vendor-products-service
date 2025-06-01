package handlers

import (
	"customer-mock-api-gateway/customer-mock-api-gateway/internal/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCartHandler(c *gin.Context) {

	cart := models.Cart{
		Id: "mockCartId",
		Items: []models.Item{
			{
				ProductId: "mockProductId",
				Name:      "mockProductName",
				Price:     500,
				Quantity:  1,
			},
			{
				ProductId: "mockProductId2",
				Name:      "mockProductName2",
				Price:     600,
				Quantity:  4,
			},
			{
				ProductId: "mockProductId3",
				Name:      "mockProductName3",
				Price:     200,
				Quantity:  10,
			},
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": cart})

}

func PostCartHandler(c *gin.Context) {

	var cartAdditions []models.CartAddition

	if err := c.ShouldBindJSON(&cartAdditions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid items data"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": cartAdditions})

}

func PutCartQuantityHandler(c *gin.Context) {

	stockId := c.Param("productId")

	if stockId != "mockProductId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var quantity models.QuantityUpdate

	if err := c.ShouldBindJSON(&quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func DeleteCartProductHandler(c *gin.Context) {

	productId := c.Param("productId")

	if productId != "mockProductId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deletedId": productId})

}

func PostCheckoutHandler(c *gin.Context) {

	var checkoutReq models.CheckoutRequest

	if err := c.ShouldBindJSON(&checkoutReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid items data"})
		return
	}

	checkoutRes := models.CheckoutResponse{
		CheckoutId: "mockCheckoutId",
		PaymentUrl: "https://external-gateway.com/pay?checkoutId=...",
	}

	c.JSON(http.StatusCreated, gin.H{"data": checkoutRes})

}

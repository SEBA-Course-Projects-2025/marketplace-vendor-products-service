package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"vendor-mock-api-gateway/vendor-mock-api-gateway/internal/models"
)

func GetOneOrderHandler(c *gin.Context) {

	orderId := c.Param("orderId")

	order := models.Order{
		Id:         "mockOrderId",
		CustomerId: "mockCustomerId",
		VendorId:   "mockVendorId",
		Items: []string{
			"mockItem1",
			"mockItem2",
		},
		TotalPrice:       450,
		Status:           "Shipping",
		Date:             time.Now(),
		VendorConfStatus: "Approved",
	}

	if orderId != order.Id {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": order})

}

func GetOrdersHandler(c *gin.Context) {

	orders := []models.Order{
		{
			Id:         "mockOrderId",
			CustomerId: "mockCustomerId",
			VendorId:   "mockVendorId",
			Items: []string{
				"mockItem1",
				"mockItem2",
			},
			TotalPrice:       450,
			Status:           "Shipping",
			Date:             time.Now(),
			VendorConfStatus: "Approved",
		},
		{
			Id:         "mockOrderId2",
			CustomerId: "mockCustomerId2",
			VendorId:   "mockVendorId",
			Items: []string{
				"mockItem3",
				"mockItem4",
			},
			TotalPrice:       100,
			Status:           "Shipping",
			Date:             time.Now(),
			VendorConfStatus: "Approved",
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": orders})

}

func PutOrderHandler(c *gin.Context) {

	orderId := c.Param("orderId")

	if orderId != "mockOrderId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var status models.VendorConfStatusUpdate

	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

package handlers

import (
	"dev-vendor/product-service/internal/stocks/application/services"
	"dev-vendor/product-service/internal/stocks/dtos"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func (h *StockHandler) PutStockProductHandler(c *gin.Context) {

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid vendorId"})
		return
	}

	stockIdStr := c.Param("stockId")
	productIdStr := c.Param("productId")

	stockId, err := uuid.Parse(stockIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	productId, err := uuid.Parse(productIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var stockProductReq dtos.PutStockProductRequest

	if err := c.ShouldBindJSON(&stockProductReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock data"})
		return
	}

	err = services.PutStockProduct(c.Request.Context(), h.StockRepo, h.ProductRepo, stockProductReq, stockId, productId, vendorId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

package handlers

import (
	"dev-vendor/product-service/internal/stocks/application/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// DeleteStockProductByIdHandler godoc
// @Summary      Delete product from stock by ID
// @Description  Deletes a specific product from a given stock for the vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        stockId     path   string true  "Stock ID (UUID)"
// @Param        productId   path   string true  "Product ID (UUID)"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/stockId/productId"
// @Failure      404 {object} map[string]interface{} "Stock product not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks/{stockId}/products/{productId} [delete]
func (h *StockHandler) DeleteStockProductByIdHandler(c *gin.Context) {

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendorId"})
		return
	}

	stockIdStr := c.Param("stockId")
	productIdStr := c.Param("productId")

	stockId, err := uuid.Parse(stockIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock UUID"})
		return
	}

	productId, err := uuid.Parse(productIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product UUID"})
		return
	}

	if err := services.DeleteStockProductById(c.Request.Context(), h.StockRepo, stockId, productId, vendorId); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

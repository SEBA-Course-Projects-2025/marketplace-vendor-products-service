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

// PutStockHandler godoc
// @Summary      Modify stock by ID
// @Description  Modify a stock by its ID for the given vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        X-Vendor-Id header string       true  "Vendor ID (UUID)"
// @Param        stockId     path   string       true  "Stock ID (UUID)"
// @Param        request     body   dtos.PutStockRequest true "Full stock data for replacement"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/stockId/stock data"
// @Failure      404 {object} map[string]interface{} "Stock not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks/{stockId} [put]
func (h *StockHandler) PutStockHandler(c *gin.Context) {

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendorId"})
		return
	}

	idStr := c.Param("stockId")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var stockReq dtos.PutStockRequest

	if err := c.ShouldBindJSON(&stockReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock data"})
		return
	}

	err = services.PutStock(c.Request.Context(), h.StockRepo, stockReq, id, vendorId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

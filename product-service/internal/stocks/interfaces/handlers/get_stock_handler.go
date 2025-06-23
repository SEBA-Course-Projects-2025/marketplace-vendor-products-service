package handlers

import (
	"dev-vendor/product-service/internal/stocks/application/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// GetStockHandler godoc
// @Summary      Get stock by ID
// @Description  Returns a single stock by its ID for the given vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        stockId     path   string true  "Stock ID (UUID)"
// @Success      200 {object} dtos.OneStockResponse
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/stockId"
// @Failure      404 {object} map[string]interface{} "Stock not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks/{stockId} [get]
func (h *StockHandler) GetStockHandler(c *gin.Context) {

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

	stock, err := services.GetStockById(c.Request.Context(), h.StockRepo, id, vendorId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stock)

}

package handlers

import (
	"dev-vendor/product-service/internal/stocks/application/services"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// PostStockHandler godoc
// @Summary      Create a new stock
// @Description  Creates a new stock for the given vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        X-Vendor-Id header string     true  "Vendor ID (UUID)"
// @Param        request     body   dtos.StockRequest true "Stock creation payload"
// @Success      201 {object} dtos.PostStockResponse
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/stock data"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks [post]
func (h *StockHandler) PostStockHandler(c *gin.Context) {

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendorId"})
		return
	}

	var stockReq dtos.StockRequest

	if err := c.ShouldBindJSON(&stockReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid stock data"})
		return
	}

	newStock, err := services.PostStock(c.Request.Context(), h.StockRepo, h.ProductRepo, h.Db, stockReq, vendorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newStock)

}

package handlers

import (
	"dev-vendor/product-service/internal/shared/metrics"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/application/services"
	"dev-vendor/product-service/internal/stocks/dtos"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// PatchStockByIdHandler godoc
// @Summary      Partially update stock by ID
// @Description  Partially updates a stock by its ID for the given vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        stockId     path   string          true  "Stock ID (UUID)"
// @Param        request     body   dtos.StockPatchRequest true "Fields to update"
// @Success      200 {object} dtos.OneStockResponse
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/stockId/stock data"
// @Failure      404 {object} map[string]interface{}  "Stock not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks/{stockId} [patch]
func (h *StockHandler) PatchStockByIdHandler(c *gin.Context) {

	ctx, span := tracer.Tracer.Start(c.Request.Context(), "PatchStockByIdHandler")
	defer span.End()

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

	var stockReq dtos.StockPatchRequest

	if err := c.ShouldBindJSON(&stockReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stock, err := services.PatchStockById(ctx, h.StockRepo, stockReq, id, vendorId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metrics.StocksUpdatedCounter.Inc()
	c.JSON(http.StatusOK, stock)
}

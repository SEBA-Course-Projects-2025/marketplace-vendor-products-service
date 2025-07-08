package handlers

import (
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/application/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// DeleteManyStockProductsHandler godoc
// @Summary      Delete multiple products from a stock
// @Description  Deletes multiple products specified by their IDs from a given stock for the vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        stockId     path   string          true  "Stock ID (UUID)"
// @Param        request     body   dtos.IdsToDelete true  "List of product IDs to delete"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/stockId/IDs of products to delete"
// @Failure      404 {object} map[string]interface{} "Stock products not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks/{stockId}/products [delete]
func (h *StockHandler) DeleteManyStockProductsHandler(c *gin.Context) {

	ctx, span := tracer.Tracer.Start(c.Request.Context(), "DeleteManyStockProductsHandler")
	defer span.End()

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendorId"})
		return
	}

	stockIdStr := c.Param("stockId")

	stockId, err := uuid.Parse(stockIdStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	var ids dtos.IdsToDelete

	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DeleteManyStockProducts(ctx, h.StockRepo, ids.Ids, stockId, vendorId); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock products not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

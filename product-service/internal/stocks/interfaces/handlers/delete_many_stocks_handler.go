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

// DeleteManyStocksHandler godoc
// @Summary      Delete multiple stocks by IDs
// @Description  Deletes multiple stocks specified by their IDs for the given vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        request     body   dtos.IdsToDelete true  "List of stock IDs to delete"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/IDs of stocks to delete"
// @Failure      404 {object} map[string]interface{} "Stock not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks [delete]
func (h *StockHandler) DeleteManyStocksHandler(c *gin.Context) {

	ctx, span := tracer.Tracer.Start(c.Request.Context(), "DeleteManyStocksHandler")
	defer span.End()

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendorId"})
		return
	}

	var ids dtos.IdsToDelete

	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.DeleteManyStocks(ctx, h.StockRepo, ids.Ids, vendorId); err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

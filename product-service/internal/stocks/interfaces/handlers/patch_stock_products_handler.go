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

// PatchStockProductsHandler godoc
// @Summary      Partially update multiple products in a stock
// @Description  Partially updates multiple products within a given stock for the vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        X-Vendor-Id header string                            true  "Vendor ID (UUID)"
// @Param        stockId     path   string                            true  "Stock ID (UUID)"
// @Param        request     body   []dtos.PatchStockManyProductsRequest true "List of products with fields to update"
// @Success      200 {array}  dtos.StockProductInfo
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/stockId/stock product data"
// @Failure      404 {object} map[string]interface{} "Stock products not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks/{stockId}/products [patch]
func (h *StockHandler) PatchStockProductsHandler(c *gin.Context) {

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

	var stockProductReq []dtos.PatchStockManyProductsRequest

	if err := c.ShouldBindJSON(&stockProductReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stockProducts, err := services.PatchStockProducts(c.Request.Context(), h.StockRepo, h.ProductRepo, h.Db, stockProductReq, stockId, vendorId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock products not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stockProducts)
}

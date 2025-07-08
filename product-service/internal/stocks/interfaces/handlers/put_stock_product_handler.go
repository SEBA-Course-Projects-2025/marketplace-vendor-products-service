package handlers

import (
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/application/services"
	"dev-vendor/product-service/internal/stocks/dtos"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// PutStockProductHandler godoc
// @Summary      Modify a product in a stock by IDs
// @Description  Modifies a specific product within a given stock for the vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        stockId     path   string           true  "Stock ID (UUID)"
// @Param        productId   path   string           true  "Product ID (UUID)"
// @Param        request     body   dtos.PutStockProductRequest true "Full product data for replacement"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/stockId/productId/stock product data"
// @Failure      404 {object} map[string]interface{} "Stock product not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks/{stockId}/products/{productId} [put]
func (h *StockHandler) PutStockProductHandler(c *gin.Context) {

	ctx, span := tracer.Tracer.Start(c.Request.Context(), "PutStockProductHandler")
	defer span.End()

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

	err = services.PutStockProduct(ctx, h.StockRepo, h.ProductRepo, h.EventRepo, h.Db, stockProductReq, stockId, productId, vendorId)
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

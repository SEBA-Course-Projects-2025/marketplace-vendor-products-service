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

// PatchStockProductByIdHandler godoc
// @Summary      Partially update a product in a stock by IDs
// @Description  Partially updates the details of a specific product within a given stock for the vendor.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        stockId     path   string                      true  "Stock ID (UUID)"
// @Param        productId   path   string                      true  "Product ID (UUID)"
// @Param        request     body   dtos.PatchStockProductRequest true "Fields to update in the stock product"
// @Success      200 {object} dtos.StockProductInfo
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/stockId/productId/stock product data"
// @Failure      404 {object} map[string]interface{} "Stock product not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks/{stockId}/products/{productId} [patch]
func (h *StockHandler) PatchStockProductByIdHandler(c *gin.Context) {

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

	var stockProductReq dtos.PatchStockProductRequest

	if err := c.ShouldBindJSON(&stockProductReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stock, err := services.PatchStockProductById(c.Request.Context(), h.StockRepo, h.ProductRepo, h.Db, stockProductReq, stockId, productId, vendorId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Stock product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stock)
}

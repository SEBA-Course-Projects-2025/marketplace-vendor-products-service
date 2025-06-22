package handlers

import (
	"dev-vendor/product-service/internal/products/application/services"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// GetProductHandler godoc
// @Summary      Get product by ID
// @Description  Returns a single product by its ID for the given vendor.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        X-Vendor-Id header string true  "Vendor ID (UUID)"
// @Param        productId   path   string true  "Product ID (UUID)"
// @Success      200 {object} dtos.OneProductResponse
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/productId"
// @Failure      404 {object} map[string]interface{} "Product not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /products/{productId} [get]
func (h *ProductHandler) GetProductHandler(c *gin.Context) {

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendorId"})
		return
	}

	idStr := c.Param("productId")

	id, err := uuid.Parse(idStr)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	product, err := services.GetProductById(c.Request.Context(), h.ProductRepo, id, vendorId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)

}

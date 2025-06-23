package handlers

import (
	"dev-vendor/product-service/internal/products/application/services"
	"dev-vendor/product-service/internal/products/dtos"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// PatchProductHandler godoc
// @Summary      Update product fields by ID
// @Description  Partially updates a product by its ID for the given vendor.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        productId   path   string true  "Product ID (UUID)"
// @Param        request     body   dtos.ProductPatchRequest true "Fields to update"
// @Success      200 {object} dtos.OneProductResponse
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/productId/product data"
// @Failure      404 {object} map[string]interface{} "Product not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /products/{productId} [patch]
func (h *ProductHandler) PatchProductHandler(c *gin.Context) {

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

	var productReq dtos.ProductPatchRequest

	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := services.PatchProduct(c.Request.Context(), h.ProductRepo, id, productReq, vendorId)

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

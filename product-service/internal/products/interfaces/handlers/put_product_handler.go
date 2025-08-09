package handlers

import (
	"dev-vendor/product-service/internal/products/application/services"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/metrics"
	"dev-vendor/product-service/internal/shared/tracer"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

// PutProductHandler godoc
// @Summary      Update product by ID
// @Description  Updates an existing product by its ID for the given vendor.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        productId   path   string true  "Product ID (UUID)"
// @Param        request     body   dtos.ProductRequest true "Full product data"
// @Success      200 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/productId/product data"
// @Failure      404 {object} map[string]interface{} "Product not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /products/id/{productId} [put]
func (h *ProductHandler) PutProductHandler(c *gin.Context) {

	ctx, span := tracer.Tracer.Start(c.Request.Context(), "PutProductByIdHandler")
	defer span.End()

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

	var productReq dtos.ProductRequest

	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	err = services.PutProduct(ctx, h.ProductRepo, h.EventRepo, h.Db, id, productReq, vendorId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metrics.ProductsUpdatedCounter.Inc()
	c.JSON(http.StatusOK, gin.H{})

}

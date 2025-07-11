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

// PatchProductBySlugHandler godoc
// @Summary      Update product fields by ID
// @Description  Partially updates a product by its ID for the given vendor.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        slug   path   string true  "Product slug (unique identifier in URL)"
// @Param        request     body   dtos.ProductPatchRequest true "Fields to update"
// @Success      200 {object} dtos.OneProductResponse
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/slug/product data"
// @Failure      404 {object} map[string]interface{} "Product not found"
// @Failure      500 {object} map[string]interface{}
// @Router       /products/slug/{slug} [patch]
func (h *ProductHandler) PatchProductBySlugHandler(c *gin.Context) {

	ctx, span := tracer.Tracer.Start(c.Request.Context(), "PatchProductBySlugHandler")
	defer span.End()

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendorId"})
		return
	}

	slug := c.Param("slug")

	var productReq dtos.ProductPatchRequest

	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	product, err := services.PatchProductBySlug(ctx, h.ProductRepo, h.EventRepo, h.Db, slug, productReq, vendorId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metrics.ProductsUpdatedCounter.Inc()
	c.JSON(http.StatusOK, product)

}

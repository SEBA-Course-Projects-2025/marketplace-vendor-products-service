package handlers

import (
	"dev-vendor/product-service/internal/products/application/services"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/metrics"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// PostProductHandler godoc
// @Summary      Create a new product
// @Description  Creates a new product for the given vendor.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        request     body   dtos.ProductRequest true "Product creation payload"
// @Success      201 {object} dtos.OneProductResponse
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/product data"
// @Failure      500 {object} map[string]interface{}
// @Router       /products [post]
func (h *ProductHandler) PostProductHandler(c *gin.Context) {

	ctx, span := tracer.Tracer.Start(c.Request.Context(), "PostProductHandler")
	defer span.End()

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendorId"})
		return
	}

	var product dtos.ProductRequest

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	newProduct, err := services.PostProduct(ctx, h.ProductRepo, h.EventRepo, h.Db, product, vendorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	metrics.ProductsAddedCounter.Inc()
	c.JSON(http.StatusCreated, newProduct)

}

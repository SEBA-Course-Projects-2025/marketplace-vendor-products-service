package handlers

import (
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/application/services"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// GetAllStockProductsHandler godoc
// @Summary      Get all stock products
// @Description  Returns a paginated list of stock products for the given vendor with optional sorting.
// @Tags         stocks
// @Accept       json
// @Produce      json
// @Param        Authorization header string true "Bearer access token"
// @Param        page      query  int    false "Page number (default is 1)"
// @Param        size      query  int    false "Page size (default is 10)"
// @Param        offset    query  int    false "Custom offset (overrides page if provided)"
// @Param        limit     query  int    false "Custom limit (overrides size if provided)"
// @Param        sortBy    query  string false "Field to sort by (default is name)"
// @Param        sortOrder query  string false "Sort order: asc or desc (default is asc)"
// @Success      200 {array} dtos.StockProductsResponseDto
// @Failure      400 {object} map[string]interface{} "Invalid vendorId/page/page size/limit/offset"
// @Failure      500 {object} map[string]interface{}
// @Router       /stocks/:stockId/products [get]
func (h *StockHandler) GetAllStockProductsHandler(c *gin.Context) {

	ctx, span := tracer.Tracer.Start(c.Request.Context(), "GetAllStockProductsHandler")
	defer span.End()

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid vendorId"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}

	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	offset, err := strconv.Atoi(c.DefaultQuery("offset", "-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid offset"})
		return
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "-1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	if offset < 0 || limit <= 0 {
		offset = (page - 1) * size
		limit = size
	}

	sortBy := c.DefaultQuery("sortBy", "name")
	sortOrder := c.DefaultQuery("sortOrder", "asc")

	queryParams := dtos.StockProductsQueryParams{
		Limit:     limit,
		Offset:    offset,
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}

	stockProducts, err := services.GetAllStockProducts(ctx, h.StockRepo, queryParams, vendorId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stockProducts)

}

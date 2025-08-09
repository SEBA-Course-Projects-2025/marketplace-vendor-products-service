package handlers

import (
	"dev-vendor/product-service/internal/products/application/services"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

// GetAllProductsHandler godoc
// @Summary      Get all products
// @Description  Returns a paginated list of products for the given vendor, with filtering, searching and sorting options.
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page       query     int     false  "Page number"      default(1)   minimum(1)
// @Param        size       query     int     false  "Page size"        default(15)  minimum(1)
// @Param        offset     query     int     false  "Offset"           default(-1)
// @Param        limit      query     int     false  "Limit"            default(-1)
// @Param        category   query     string  false  "Category filter"
// @Param        minPrice   query     number  false  "Minimum price"    default(0)   minimum(0)
// @Param        maxPrice   query     number  false  "Maximum price"    default(0)   minimum(0)
// @Param        search     query     string  false  "Search term"
// @Param        sortBy     query     string  false  "Sort by field"    default(name)
// @Param        sortOrder  query     string  false  "Sort order"       default(asc) Enums(asc,desc)
// @Param        Authorization header string true "Bearer access token"
// @Success      200        {array}   dtos.GetProductsResponse
// @Failure      400        {object}  map[string]interface{} "Invalid vendorId/page/page size/offset/limit/price"
// @Failure      500        {object}  map[string]interface{}
// @Router       /products [get]
func (h *ProductHandler) GetAllProductsHandler(c *gin.Context) {

	ctx, span := tracer.Tracer.Start(c.Request.Context(), "GetAllProductsHandler")
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

	size, err := strconv.Atoi(c.DefaultQuery("size", "15"))
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

	category := c.Query("category")

	minPrice, err := strconv.ParseFloat(c.DefaultQuery("minPrice", "0"), 64)
	if err != nil || minPrice < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid minimum price"})
		return
	}

	maxPrice, err := strconv.ParseFloat(c.DefaultQuery("maxPrice", "0"), 64)
	if err != nil || maxPrice < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid maximum price"})
		return
	}

	if minPrice > maxPrice && maxPrice > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid prices range"})
		return
	}

	search := c.Query("search")

	sortBy := c.DefaultQuery("sortBy", "name")
	sortOrder := c.DefaultQuery("sortOrder", "asc")

	queryParams := dtos.ProductQueryParams{
		Limit:     limit,
		Offset:    offset,
		Category:  category,
		MinPrice:  minPrice,
		MaxPrice:  maxPrice,
		SortBy:    sortBy,
		SortOrder: sortOrder,
		Search:    search,
	}

	products, err := services.GetAllProducts(ctx, h.ProductRepo, queryParams, vendorId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)

}

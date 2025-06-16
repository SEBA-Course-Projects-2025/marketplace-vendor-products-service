package handlers

import (
	"dev-vendor/product-service/internal/products/application/services"
	"dev-vendor/product-service/internal/products/dtos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func (h *ProductHandler) GetAllProductsHandler(c *gin.Context) {

	v, _ := c.Get("vendorId")
	vendorId, ok := v.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid vendorId"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	size, _ := strconv.Atoi(c.DefaultQuery("size", "15"))

	offset, _ := strconv.Atoi(c.Query("offset"))

	limit, _ := strconv.Atoi(c.Query("limit"))

	if offset < 0 || limit <= 0 {
		offset = (page - 1) * size
		limit = size
	}

	category := c.Query("category")

	minPrice, _ := strconv.ParseFloat(c.DefaultQuery("minPrice", "0"), 64)

	maxPrice, _ := strconv.ParseFloat(c.DefaultQuery("maxPrice", "0"), 64)

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

	products, err := services.GetAllProducts(c.Request.Context(), h.ProductRepo, queryParams, vendorId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)

}

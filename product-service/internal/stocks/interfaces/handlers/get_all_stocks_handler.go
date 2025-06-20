package handlers

import (
	"dev-vendor/product-service/internal/stocks/application/services"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

func (h *StockHandler) GetAllStocksHandler(c *gin.Context) {

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

	locationId := c.Query("location_id")

	sortBy := c.DefaultQuery("sortBy", "date_supplied")
	sortOrder := c.DefaultQuery("sortOrder", "asc")

	queryParams := dtos.StockQueryParams{
		Limit:      limit,
		Offset:     offset,
		LocationId: locationId,
		SortBy:     sortBy,
		SortOrder:  sortOrder,
	}

	stocks, err := services.GetAllStocks(c.Request.Context(), h.StockRepo, queryParams, vendorId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stocks)

}

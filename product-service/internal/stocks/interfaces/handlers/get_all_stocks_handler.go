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

	locationId := c.Query("location_id")

	sortBy := c.DefaultQuery("sortBy", "name")
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

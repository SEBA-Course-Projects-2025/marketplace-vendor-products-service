package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

func VendorIdMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		vendorIdStr := c.GetHeader("X-Vendor-Id")

		vendorId, err := uuid.Parse(vendorIdStr)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "X-Vendor-Id header is required"})
			return
		}

		c.Set("vendorId", vendorId)
		c.Next()

	}
}

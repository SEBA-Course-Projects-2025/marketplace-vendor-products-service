package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthenticationMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader != "mock-jwt-token" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}

}

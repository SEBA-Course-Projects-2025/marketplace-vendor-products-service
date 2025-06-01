package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vendor-mock-api-gateway/vendor-mock-api-gateway/internal/models"
)

func GetProfileHandler(c *gin.Context) {

	id := c.Param("vendorId")

	profile := models.Profile{
		Id:          "mockVendorId",
		Name:        "mock-vendor",
		Email:       "vendor@example.com",
		Description: "mock description information about the vendor",
		Logo:        "https://example.com/vendorLogo.png",
		Address:     "123 Vendor St, Vendor City",
		Website:     "https://mock-vendor.com",
		CatalogId:   "de305d54-75b4-431b-adb2-eb6b9e546014",
	}

	if profile.Id == id {
		c.JSON(http.StatusOK, gin.H{
			"data": profile,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{})

}

func PutProfileHandler(c *gin.Context) {

	id := c.Param("vendorId")

	var profile models.Profile

	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid profile data"})
		return
	}

	if id != "mockVendorId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

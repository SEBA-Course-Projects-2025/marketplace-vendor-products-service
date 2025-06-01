package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mock-api-gateway/mock-api-gateways/vendor-mock-api-gateway/internal/models"
	"net/http"
)

func GetOneProductHandler(c *gin.Context) {

	productId := c.Param("productId")

	product := models.Product{
		Id:          "mockProductId",
		Name:        "mock-product",
		Description: "mock description information about the  product",
		Price:       140,
		Images: []string{
			"https://example.com/productImage01.png",
			"https://example.com/productImage02.png",
		},
		Availability: "InStock",
		Category:     "mock-category",
		Tags: []string{
			"mock-tag1",
			"mock-tag2",
		},
		VendorId: "mockVendorId",
	}

	if product.Id != productId {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func GetProductsHandler(c *gin.Context) {

	products := []models.Product{
		{
			Id:          "mockProductId",
			Name:        "mock-product",
			Description: "mock description information about the product",
			Price:       140,
			Images: []string{
				"https://example.com/productImage01.png",
				"https://example.com/productImage02.png",
			},
			Availability: "InStock",
			Category:     "mock-category",
			Tags: []string{
				"mock-tag1",
				"mock-tag2",
			},
			VendorId: "mockVendorId",
		},
		{
			Id:          "mockProductId2",
			Name:        "mock-product2",
			Description: "mock description information about the product2",
			Price:       200.65,
			Images: []string{
				"https://example.com/productImage03.png",
				"https://example.com/productImage04.png",
			},
			Availability: "OutOfStock",
			Category:     "mock-category2",
			Tags: []string{
				"mock-tag1",
				"mock-tag4",
			},
			VendorId: "mockVendorId2",
		},
		{
			Id:          "mockProductId3",
			Name:        "mock-product3",
			Description: "mock description information about the product3",
			Price:       200.65,
			Images: []string{
				"https://example.com/productImage05.png",
				"https://example.com/productImage06.png",
			},
			Availability: "InStock",
			Category:     "mock-category1",
			Tags: []string{
				"mock-tag2",
				"mock-tag3",
			},
			VendorId: "mockVendorId2",
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": products})

}

func PostProductsHandler(c *gin.Context) {

	var products []models.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid products data"})
		return
	}

	for i := range products {
		products[i].Id = fmt.Sprintf("mock-id-%d", i+1)
	}

	c.JSON(http.StatusCreated, gin.H{"data": products})

}

func PutProductHandler(c *gin.Context) {

	productId := c.Param("productId")

	if productId != "mockProductId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})

}

func PatchOneProductHandler(c *gin.Context) {

	productId := c.Param("productId")

	if productId != "mockProductId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var product models.Product

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	product.Id = productId

	c.JSON(http.StatusOK, gin.H{"data": product})

}

func PatchProductsHandler(c *gin.Context) {

	var products []models.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": products})

}

func DeleteOneProductHandler(c *gin.Context) {

	productId := c.Param("productId")

	if productId != "mockProductId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deletedId": productId})

}

func DeleteProductsHandler(c *gin.Context) {

	var ids models.IdsToDelete

	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"deletedIds": ids.Ids})

}

func GetOneStockProductHandler(c *gin.Context) {

	stockId := c.Param("stockId")

	stockProduct := models.StockProduct{
		Id:           "mockStockId",
		VendorId:     "mockVendorId",
		Name:         "mock-product",
		Image:        "https://example.com/productImage01.png",
		Availability: "InStock",
		Quantity:     10,
	}

	if stockProduct.Id != stockId {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": stockProduct})

}

func GetStockProductsHandler(c *gin.Context) {

	stockProducts := []models.StockProduct{
		{
			Id:           "mockStockId",
			VendorId:     "mockVendorId",
			Name:         "mock-product",
			Image:        "https://example.com/productImage01.png",
			Availability: "InStock",
			Quantity:     10,
		},
		{
			Id:           "mockStockId2",
			VendorId:     "mockVendorId2",
			Name:         "mock-product2",
			Image:        "https://example.com/productImage03.png",
			Availability: "OutOfStock",
			Quantity:     1,
		},
		{
			Id:           "mockStockId3",
			VendorId:     "mockVendorId2",
			Name:         "mock-product3",
			Image:        "https://example.com/productImage05.png",
			Availability: "InStock",
			Quantity:     4,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": stockProducts})

}

func PutStockProductHandler(c *gin.Context) {

	stockId := c.Param("stockId")

	if stockId != "mockStockId" {
		c.JSON(http.StatusNotFound, gin.H{})
		return
	}

	var quantity models.QuantityUpdate

	if err := c.ShouldBindJSON(&quantity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quantity data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

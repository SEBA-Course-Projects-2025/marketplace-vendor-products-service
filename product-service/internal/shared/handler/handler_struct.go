package handlers

import (
	productDomain "dev-vendor/product-service/internal/products/domain"
	stockDomain "dev-vendor/product-service/internal/stocks/domain"
)

type Handler struct {

	ProductRepo productDomain.ProductRepository

	StockRepo stockDomain.StockRepository
}

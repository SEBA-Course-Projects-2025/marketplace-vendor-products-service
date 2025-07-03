package handlers

import (
	eventDomain "dev-vendor/product-service/internal/event/domain"
	productDomain "dev-vendor/product-service/internal/products/domain"
	stockDomain "dev-vendor/product-service/internal/stocks/domain"
	"gorm.io/gorm"
)

type Handler struct {
	ProductRepo productDomain.ProductRepository

	StockRepo stockDomain.StockRepository

	EventRepo eventDomain.EventRepository

	Db *gorm.DB
}

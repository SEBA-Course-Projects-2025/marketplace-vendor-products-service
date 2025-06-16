package domain

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

type StockRepository interface {
	FindById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) (*models.Stock, error)
	FindAll(ctx context.Context, params dtos.StockQueryParams, vendorId uuid.UUID) (*[]models.Stock, error)
	Create(ctx context.Context, newStock *models.Stock, vendorId uuid.UUID) (*models.Stock, error)
	UpdateStock(ctx context.Context, updatedStock *models.Stock) error
	UpdateStockProduct(ctx context.Context, updatedStockProduct *models.StocksProduct) error
	PatchStockId(ctx context.Context, modifiedStock *models.Stock) (*models.Stock, error)
	PatchStockProductId(ctx context.Context, modifiedStockProduct *models.StocksProduct) (*models.StocksProduct, error)
	PatchStockProducts(ctx context.Context, modifiedStockProducts []models.StocksProduct) ([]models.StocksProduct, error)
	DeleteStockById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) error
	DeleteStockProductById(ctx context.Context, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) error
	DeleteManyStocks(ctx context.Context, ids []uuid.UUID, vendorId uuid.UUID) error
	DeleteManyStockProducts(ctx context.Context, ids []uuid.UUID, stockId uuid.UUID, vendorId uuid.UUID) error
	CheckProduct(ctx context.Context, productId uuid.UUID, vendorId uuid.UUID) error
	CheckLocation(ctx context.Context, locationId uuid.UUID) (*models.StocksLocation, error)
}

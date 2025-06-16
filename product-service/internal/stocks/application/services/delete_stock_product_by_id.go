package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
)

func DeleteStockProductById(ctx context.Context, stockRepo domain.StockRepository, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) error {

	return stockRepo.DeleteStockProductById(ctx, stockId, productId, vendorId)

}

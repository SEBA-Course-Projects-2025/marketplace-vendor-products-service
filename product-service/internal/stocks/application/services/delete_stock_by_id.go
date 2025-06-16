package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
)

func DeleteStockById(ctx context.Context, stockRepo domain.StockRepository, id uuid.UUID, vendorId uuid.UUID) error {

	return stockRepo.DeleteStockById(ctx, id, vendorId)

}

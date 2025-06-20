package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
)

func DeleteManyStocks(ctx context.Context, stockRepo domain.StockRepository, ids []uuid.UUID, vendorId uuid.UUID) error {

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		return txRepo.DeleteManyStocks(ctx, ids, vendorId)
	})

}

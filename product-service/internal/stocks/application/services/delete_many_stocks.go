package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
)

func DeleteManyStocks(ctx context.Context, stockRepo domain.StockRepository, ids []uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteManyStocks")
	defer span.End()

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		return txRepo.DeleteManyStocks(ctx, ids, vendorId)
	})

}

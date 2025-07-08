package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
)

func DeleteManyStockProducts(ctx context.Context, stockRepo domain.StockRepository, ids []uuid.UUID, stockId uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteManyStockProducts")
	defer span.End()

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		return txRepo.DeleteManyStockProducts(ctx, ids, stockId, vendorId)
	})

}

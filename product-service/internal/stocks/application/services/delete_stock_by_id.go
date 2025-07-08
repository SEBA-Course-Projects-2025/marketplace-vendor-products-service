package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
)

func DeleteStockById(ctx context.Context, stockRepo domain.StockRepository, id uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteStockById")
	defer span.End()

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		return txRepo.DeleteStockById(ctx, id, vendorId)
	})

}

package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
)

func DeleteStockProductById(ctx context.Context, stockRepo domain.StockRepository, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteStockProductById")
	defer span.End()

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		return txRepo.DeleteStockProductById(ctx, stockId, productId, vendorId)
	})

}

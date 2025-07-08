package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func PutStock(ctx context.Context, stockRepo domain.StockRepository, stockReq dtos.PutStockRequest, stockId uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "PutStock")
	defer span.End()

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		location, err := txRepo.CheckLocation(ctx, stockReq.LocationId)

		if err != nil {
			return err
		}

		existingStock, err := txRepo.FindById(ctx, stockId)

		if err != nil {
			return err
		}

		existingStock = dtos.UpdateStockWithDto(existingStock, stockReq, location)

		return txRepo.UpdateStock(ctx, existingStock)
	})

}

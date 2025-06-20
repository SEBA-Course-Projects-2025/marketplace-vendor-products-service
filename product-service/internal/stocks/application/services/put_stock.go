package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func PutStock(ctx context.Context, stockRepo domain.StockRepository, stockReq dtos.PutStockRequest, stockId uuid.UUID, vendorId uuid.UUID) error {

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		location, err := txRepo.CheckLocation(ctx, stockReq.LocationId)

		if err != nil {
			return err
		}

		existingStock, err := txRepo.FindById(ctx, stockId, vendorId)

		if err != nil {
			return err
		}

		existingStock = dtos.UpdateStockWithDto(existingStock, stockReq, location)

		return txRepo.UpdateStock(ctx, existingStock)
	})

}

package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func PatchStockById(ctx context.Context, stockRepo domain.StockRepository, stockReq dtos.StockPatchRequest, stockId uuid.UUID, vendorId uuid.UUID) (dtos.OneStockResponse, error) {

	var stockResponse dtos.OneStockResponse

	err := stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		var location *models.StocksLocation

		if stockReq.LocationId != nil && *stockReq.LocationId != uuid.Nil {
			var err error
			location, err = txRepo.CheckLocation(ctx, *stockReq.LocationId)

			if err != nil {
				return err
			}
		}

		existingStock, err := txRepo.FindById(ctx, stockId, vendorId)

		if err != nil {
			return err
		}

		existingStock = dtos.ModifyStockWithDto(existingStock, stockReq, location)

		updatedStock, err := txRepo.PatchStockId(ctx, existingStock)
		if err != nil {
			return err
		}

		stockResponse = dtos.StockToDto(updatedStock)

		return nil

	})

	if err != nil {
		return dtos.OneStockResponse{}, err
	}

	return stockResponse, nil

}

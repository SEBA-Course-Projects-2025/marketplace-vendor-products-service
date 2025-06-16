package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func PatchStockById(ctx context.Context, stockRepo domain.StockRepository, stockReq dtos.StockPatchRequest, stockId uuid.UUID, vendorId uuid.UUID) (dtos.OneStockResponse, error) {

	var location *models.StocksLocation

	if stockReq.LocationId != nil && *stockReq.LocationId != uuid.Nil {
		var err error
		location, err = stockRepo.CheckLocation(ctx, *stockReq.LocationId)

		if err != nil {
			return dtos.OneStockResponse{}, err
		}
	}

	existingStock, err := stockRepo.FindById(ctx, stockId, vendorId)

	if err != nil {
		return dtos.OneStockResponse{}, err
	}

	existingStock = dtos.ModifyStockWithDto(existingStock, stockReq, location)

	updatedStock, err := stockRepo.PatchStockId(ctx, existingStock)

	return dtos.StockToDto(updatedStock), nil

}

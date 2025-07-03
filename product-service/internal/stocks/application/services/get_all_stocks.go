package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func GetAllStocks(ctx context.Context, stockRepo domain.StockRepository, params dtos.StockQueryParams, vendorId uuid.UUID) ([]dtos.GetStocksResponse, error) {

	stocks, err := stockRepo.FindAll(ctx, params, vendorId)

	if err != nil {
		return nil, err
	}

	return dtos.StocksToDto(stocks), nil

}

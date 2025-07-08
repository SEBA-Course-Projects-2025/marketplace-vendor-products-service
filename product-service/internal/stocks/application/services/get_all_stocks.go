package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func GetAllStocks(ctx context.Context, stockRepo domain.StockRepository, params dtos.StockQueryParams, vendorId uuid.UUID) ([]dtos.GetStocksResponse, error) {

	ctx, span := tracer.Tracer.Start(ctx, "GetAllStocks")
	defer span.End()

	stocks, err := stockRepo.FindAll(ctx, params, vendorId)

	if err != nil {
		return nil, err
	}

	return dtos.StocksToDto(stocks), nil

}

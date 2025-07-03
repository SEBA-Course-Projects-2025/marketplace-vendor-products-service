package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func GetAllStockProducts(ctx context.Context, stockRepo domain.StockRepository, params dtos.StockProductsQueryParams, vendorId uuid.UUID) ([]dtos.StockProductsResponseDto, error) {

	stockProducts, err := stockRepo.FindAllStockProducts(ctx, params, vendorId)

	if err != nil {
		return nil, err
	}

	return dtos.StockProductsToDto(stockProducts), nil

}

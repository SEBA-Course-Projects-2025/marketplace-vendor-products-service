package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func GetAllStockProducts(ctx context.Context, stockRepo domain.StockRepository, params dtos.StockProductsQueryParams, vendorId uuid.UUID) ([]dtos.StockProductsResponseDto, error) {

	ctx, span := tracer.Tracer.Start(ctx, "GetAllStockProducts")
	defer span.End()

	stockProducts, err := stockRepo.FindAllStockProducts(ctx, params, vendorId)

	if err != nil {
		return nil, err
	}

	return dtos.StockProductsToDto(stockProducts), nil

}

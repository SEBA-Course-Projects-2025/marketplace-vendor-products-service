package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func GetStockById(ctx context.Context, repo domain.StockRepository, id uuid.UUID, vendorID uuid.UUID) (dtos.OneStockResponse, error) {

	ctx, span := tracer.Tracer.Start(ctx, "GetStockById")
	defer span.End()

	stock, err := repo.FindById(ctx, id)

	if err != nil {
		return dtos.OneStockResponse{}, err
	}

	return dtos.StockToDto(stock), nil

}

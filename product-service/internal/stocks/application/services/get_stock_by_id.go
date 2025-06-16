package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func GetStockById(ctx context.Context, repo domain.StockRepository, id uuid.UUID, vendorID uuid.UUID) (dtos.OneStockResponse, error) {

	stock, err := repo.FindById(ctx, id, vendorID)

	if err != nil {
		return dtos.OneStockResponse{}, err
	}

	return dtos.StockToDto(stock), nil

}

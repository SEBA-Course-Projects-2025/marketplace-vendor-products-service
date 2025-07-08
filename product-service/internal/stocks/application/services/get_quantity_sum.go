package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
)

func GetQuantitySum(ctx context.Context, stockRepo domain.StockRepository, productId uuid.UUID) (int, error) {

	ctx, span := tracer.Tracer.Start(ctx, "GetQuantitySum")
	defer span.End()

	var quantitySum = 0

	productsStocks, err := stockRepo.FindProductStocksQuantities(ctx, productId)

	if err != nil {
		return -1, err
	}

	for _, stockProduct := range productsStocks {
		quantitySum += stockProduct.Quantity
	}

	return quantitySum, nil

}

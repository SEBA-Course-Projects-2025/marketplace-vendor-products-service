package services

import (
	"context"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
)

func GetQuantitySum(ctx context.Context, stockRepo domain.StockRepository, productId uuid.UUID, vendorId uuid.UUID) (int, error) {

	var quantitySum = 0

	productsStocks, err := stockRepo.FindProductStocksQuantities(ctx, productId, vendorId)

	if err != nil {
		return -1, err
	}

	for _, stockProduct := range productsStocks {
		quantitySum += stockProduct.Quantity
	}

	return quantitySum, nil

}

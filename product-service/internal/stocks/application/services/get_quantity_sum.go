package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetQuantitySum(ctx context.Context, stockRepo domain.StockRepository, productId uuid.UUID) (int, error) {

	logrus.WithFields(logrus.Fields{
		"productId": productId,
	}).Info("Starting GetQuantitySum application service")

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

	logrus.WithFields(logrus.Fields{
		"productId": productId,
	}).Info("Successfully get quantity sum of product from stocks")

	return quantitySum, nil

}

package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetStockById(ctx context.Context, repo domain.StockRepository, id uuid.UUID, vendorID uuid.UUID) (dtos.OneStockResponse, error) {

	logrus.WithFields(logrus.Fields{
		"stockId": id,
	}).Info("Starting GetStockById application service")

	ctx, span := tracer.Tracer.Start(ctx, "GetStockById")
	defer span.End()

	stock, err := repo.FindById(ctx, id)

	if err != nil {
		return dtos.OneStockResponse{}, err
	}

	logrus.WithFields(logrus.Fields{
		"productId": id,
	}).Info("Successfully get stock by id")

	return dtos.StockToDto(stock), nil

}

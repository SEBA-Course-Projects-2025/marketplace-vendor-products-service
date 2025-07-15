package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllStocks(ctx context.Context, stockRepo domain.StockRepository, params dtos.StockQueryParams, vendorId uuid.UUID) ([]dtos.GetStocksResponse, error) {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Starting GetAllStocks application service")

	ctx, span := tracer.Tracer.Start(ctx, "GetAllStocks")
	defer span.End()

	stocks, err := stockRepo.FindAll(ctx, params, vendorId)

	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Successfully get paginated list of stocks")

	return dtos.StocksToDto(stocks), nil

}

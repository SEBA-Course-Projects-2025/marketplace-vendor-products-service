package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllStockProducts(ctx context.Context, stockRepo domain.StockRepository, params dtos.StockProductsQueryParams, vendorId uuid.UUID) ([]dtos.StockProductsResponseDto, error) {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Starting GetAllStockProducts application service")

	ctx, span := tracer.Tracer.Start(ctx, "GetAllStockProducts")
	defer span.End()

	stockProducts, err := stockRepo.FindAllStockProducts(ctx, params, vendorId)

	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Successfully get paginated list of stock products")

	return dtos.StockProductsToDto(stockProducts), nil

}

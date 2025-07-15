package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func DeleteManyStocks(ctx context.Context, stockRepo domain.StockRepository, ids []uuid.UUID, vendorId uuid.UUID) error {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
		"ids":      ids,
	}).Info("Starting DeleteManyStocks application service")

	ctx, span := tracer.Tracer.Start(ctx, "DeleteManyStocks")
	defer span.End()

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		logrus.WithFields(logrus.Fields{
			"vendorId": vendorId,
			"ids":      ids,
		}).Info("Successfully deleted stocks")
		return txRepo.DeleteManyStocks(ctx, ids, vendorId)
	})

}

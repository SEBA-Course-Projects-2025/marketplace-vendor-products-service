package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func DeleteStockById(ctx context.Context, stockRepo domain.StockRepository, id uuid.UUID, vendorId uuid.UUID) error {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
		"stockId":  id,
	}).Info("Starting DeleteStockById application service")

	ctx, span := tracer.Tracer.Start(ctx, "DeleteStockById")
	defer span.End()

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		logrus.WithFields(logrus.Fields{
			"vendorId": vendorId,
			"stockId":  id,
		}).Info("Successfully deleted stock by id")
		return txRepo.DeleteStockById(ctx, id, vendorId)
	})

}

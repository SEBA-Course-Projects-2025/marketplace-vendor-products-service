package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func DeleteManyStockProducts(ctx context.Context, stockRepo domain.StockRepository, ids []uuid.UUID, stockId uuid.UUID, vendorId uuid.UUID) error {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
		"stockId":  stockId,
		"ids":      ids,
	}).Info("Starting DeleteManyStockProducts application service")

	ctx, span := tracer.Tracer.Start(ctx, "DeleteManyStockProducts")
	defer span.End()

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		logrus.WithFields(logrus.Fields{
			"vendorId": vendorId,
			"stockId":  stockId,
			"ids":      ids,
		}).Info("Successfully deleted stock products")
		return txRepo.DeleteManyStockProducts(ctx, ids, stockId, vendorId)
	})

}

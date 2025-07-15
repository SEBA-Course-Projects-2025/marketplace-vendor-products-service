package services

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func DeleteStockProductById(ctx context.Context, stockRepo domain.StockRepository, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) error {

	logrus.WithFields(logrus.Fields{
		"vendorId":  vendorId,
		"stockId":   stockId,
		"productId": productId,
	}).Info("Starting DeleteStockProductById application service")

	ctx, span := tracer.Tracer.Start(ctx, "DeleteStockProductById")
	defer span.End()

	return stockRepo.Transaction(func(txRepo domain.StockRepository) error {
		logrus.WithFields(logrus.Fields{
			"vendorId":  vendorId,
			"stockId":   stockId,
			"productId": productId,
		}).Info("Successfully deleted stock product by id")
		return txRepo.DeleteStockProductById(ctx, stockId, productId, vendorId)
	})

}

package services

import (
	"context"
	eventDomain "dev-vendor/product-service/internal/event/domain"
	"dev-vendor/product-service/internal/products/application/services"
	productDomain "dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func PostStock(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, stockReq dtos.StockRequest, vendorId uuid.UUID) (dtos.PostStockResponse, error) {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Starting PostStock application service")

	ctx, span := tracer.Tracer.Start(ctx, "PostStock")
	defer span.End()

	var stockResponse dtos.PostStockResponse

	if err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txStockRepo := stockRepo.WithTx(tx)
		txProductRepo := productRepo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		if _, err := txStockRepo.CheckLocation(ctx, stockReq.LocationId); err != nil {
			return err
		}

		for _, product := range stockReq.Products {

			if err := txStockRepo.CheckProduct(ctx, product.ProductId, vendorId); err != nil {
				return err
			}

		}

		newStock, err := dtos.PostStockRequestToStock(stockReq, vendorId)
		if err != nil {
			return err
		}

		createdStock, err := txStockRepo.Create(ctx, newStock, vendorId)
		if err != nil {
			return err
		}

		stockResponse = dtos.PostStockToStockResponse(createdStock)

		for _, product := range stockReq.Products {

			quantitySum, err := GetQuantitySum(ctx, txStockRepo, product.ProductId)

			if err != nil {
				return err
			}

			if err := services.UpdateProductQuantity(ctx, txProductRepo, txEventRepo, product.ProductId, quantitySum, "product.catalog.events"); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return dtos.PostStockResponse{}, err
	}

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Successfully created stock")

	return stockResponse, nil

}

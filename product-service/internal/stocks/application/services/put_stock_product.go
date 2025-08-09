package services

import (
	"context"
	eventDomain "dev-vendor/product-service/internal/event/domain"
	"dev-vendor/product-service/internal/products/application/services"
	productDomain "dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func PutStockProduct(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, stockProductReq dtos.PutStockProductRequest, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) error {

	logrus.WithFields(logrus.Fields{
		"vendorId":  vendorId,
		"stockId":   stockId,
		"productId": productId,
	}).Info("Starting PutStockProduct application service")

	ctx, span := tracer.Tracer.Start(ctx, "PutStockProduct")
	defer span.End()

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txStockRepo := stockRepo.WithTx(tx)
		txProductRepo := productRepo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		if err := txStockRepo.CheckProduct(ctx, productId, vendorId); err != nil {
			return err
		}

		existingStock, err := txStockRepo.FindById(ctx, stockId)
		if err != nil {
			return err
		}

		var existingStockProduct *models.StocksProduct
		for i := range existingStock.StocksProducts {
			if existingStock.StocksProducts[i].ProductId == productId {
				existingStockProduct = &existingStock.StocksProducts[i]
				break
			}
		}

		updatedStockProduct := dtos.UpdateStockProductWithDto(existingStockProduct, stockProductReq)

		quantitySum, err := GetQuantitySum(ctx, txStockRepo, productId)

		if err != nil {
			return err
		}

		if err := services.UpdateProductQuantity(ctx, txProductRepo, txEventRepo, productId, quantitySum, "product.catalog.events"); err != nil {
			return err
		}

		logrus.WithFields(logrus.Fields{
			"vendorId":  vendorId,
			"stockId":   stockId,
			"productId": productId,
		}).Info("Successfully partially modified stock product by its stockId and productId")

		return txStockRepo.UpdateStockProduct(ctx, updatedStockProduct)

	})

}

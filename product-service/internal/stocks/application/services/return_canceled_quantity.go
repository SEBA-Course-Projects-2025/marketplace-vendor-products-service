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
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func ReturnCanceledQuantity(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, stockEventReq []dtos.CanceledOrderItemDto) error {

	logrus.Info("Starting ReturnCanceledQuantity application service")

	ctx, span := tracer.Tracer.Start(ctx, "ReturnCanceledQuantity")
	defer span.End()

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txStockRepo := stockRepo.WithTx(tx)
		txProductRepo := productRepo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		for _, stockProduct := range stockEventReq {

			existingStock, err := txStockRepo.FindById(ctx, stockProduct.StockId)
			if err != nil {
				return err
			}

			var existingStockProduct *models.StocksProduct
			for i := range existingStock.StocksProducts {
				if existingStock.StocksProducts[i].ProductId == stockProduct.ProductId {
					existingStockProduct = &existingStock.StocksProducts[i]
					break
				}
			}

			existingStockProduct.Quantity += stockProduct.Quantity

			existingStockProduct, err = txStockRepo.PatchStockProductId(ctx, existingStockProduct)

			if err != nil {
				return err
			}

			quantitySum, err := GetQuantitySum(ctx, txStockRepo, stockProduct.ProductId)

			if err != nil {
				return err
			}

			if err := services.UpdateProductQuantity(ctx, txProductRepo, txEventRepo, stockProduct.ProductId, quantitySum, "product.catalog.events"); err != nil {
				return err
			}

		}

		logrus.Info("Successfully returned cancelled product quantities")

		return nil

	})

}

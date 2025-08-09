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

func PatchStockProducts(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, stockProductReq []dtos.PatchStockManyProductsRequest, stockId uuid.UUID, vendorId uuid.UUID) ([]dtos.StockProductInfo, error) {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
		"stockId":  stockId,
	}).Info("Starting PatchStockProducts application service")

	ctx, span := tracer.Tracer.Start(ctx, "PatchStockProducts")
	defer span.End()

	var stockProductsRes []dtos.StockProductInfo

	if err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txStockRepo := stockRepo.WithTx(tx)
		txProductRepo := productRepo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		existingStock, err := txStockRepo.FindById(ctx, stockId)
		if err != nil {
			return err
		}

		var modifiedProducts []models.StocksProduct

		for _, product := range stockProductReq {

			if err := txStockRepo.CheckProduct(ctx, product.ProductId, vendorId); err != nil {
				return err
			}

			for i := range existingStock.StocksProducts {
				if existingStock.StocksProducts[i].ProductId == product.ProductId {
					updatedStockProduct := dtos.ModifyStockManyProductsWithDto(&existingStock.StocksProducts[i], product)
					modifiedProducts = append(modifiedProducts, *updatedStockProduct)
					break
				}
			}

		}

		updatedProducts, err := txStockRepo.PatchStockProducts(ctx, modifiedProducts)
		if err != nil {
			return err
		}

		for _, updProduct := range updatedProducts {
			stockProductsRes = append(stockProductsRes, dtos.StocksProductToStockProductInfo(&updProduct))
		}

		for _, product := range stockProductReq {

			quantitySum, err := GetQuantitySum(ctx, txStockRepo, product.ProductId)

			if err != nil {
				return err
			}

			if product.Quantity != nil && *product.Quantity >= 0 {
				if err := services.UpdateProductQuantity(ctx, txProductRepo, txEventRepo, product.ProductId, quantitySum, "product.catalog.events"); err != nil {
					return err
				}
			}
		}

		return nil

	}); err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
		"stockId":  stockId,
	}).Info("Successfully partially modified stock products by stockId")

	return stockProductsRes, nil

}

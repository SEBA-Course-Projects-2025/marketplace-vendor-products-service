package services

import (
	"context"
	"dev-vendor/product-service/internal/products/application/services"
	productDomain "dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PutStockProduct(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, db *gorm.DB, stockProductReq dtos.PutStockProductRequest, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) error {

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txStockRepo := stockRepo.WithTx(tx)
		txProductRepo := productRepo.WithTx(tx)

		if err := txStockRepo.CheckProduct(ctx, productId, vendorId); err != nil {
			return err
		}

		existingStock, err := txStockRepo.FindById(ctx, stockId, vendorId)
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

		quantitySum, err := GetQuantitySum(ctx, txStockRepo, productId, vendorId)

		if err != nil {
			return err
		}

		if err := services.UpdateProductQuantity(ctx, txProductRepo, productId, vendorId, quantitySum); err != nil {
			return err
		}

		return txStockRepo.UpdateStockProduct(ctx, updatedStockProduct)

	})

}

package services

import (
	"context"
	"dev-vendor/product-service/internal/products/application/services"
	productDomain "dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PatchStockProductById(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, db *gorm.DB, stockProductReq dtos.PatchStockProductRequest, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) (dtos.StockProductInfo, error) {

	var updatedStockProductResponse dtos.StockProductInfo

	if err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txStockRepo := stockRepo.WithTx(tx)
		txProductRepo := productRepo.WithTx(tx)

		if productId != uuid.Nil {

			if err := txStockRepo.CheckProduct(ctx, productId, vendorId); err != nil {
				return err
			}

		}

		if stockProductReq.Quantity == nil || *stockProductReq.Quantity < 1 {
			return errors.New("invalid product quantity")
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

		updatedStockProduct := dtos.ModifyStockProductWithDto(existingStockProduct, stockProductReq)

		updatedStockProduct, err = txStockRepo.PatchStockProductId(ctx, updatedStockProduct)

		if err != nil {
			return err
		}

		quantitySum, err := GetQuantitySum(ctx, txStockRepo, productId, vendorId)

		if err != nil {
			return err
		}

		if err := services.UpdateProductQuantity(ctx, txProductRepo, productId, vendorId, quantitySum); err != nil {
			return err
		}

		updatedStockProductResponse = dtos.StocksProductToStockProductInfo(updatedStockProduct)

		return nil

	}); err != nil {
		return dtos.StockProductInfo{}, err
	}

	return updatedStockProductResponse, nil

}

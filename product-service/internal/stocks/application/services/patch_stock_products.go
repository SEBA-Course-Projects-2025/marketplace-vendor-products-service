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

func PatchStockProducts(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, db *gorm.DB, stockProductReq []dtos.PatchStockManyProductsRequest, stockId uuid.UUID, vendorId uuid.UUID) ([]dtos.StockProductInfo, error) {

	var stockProductsRes []dtos.StockProductInfo

	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txStockRepo := stockRepo.WithTx(tx)
		txProductRepo := productRepo.WithTx(tx)

		existingStock, err := txStockRepo.FindById(ctx, stockId, vendorId)
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

			quantitySum, err := GetQuantitySum(ctx, txStockRepo, product.ProductId, vendorId)

			if err != nil {
				return err
			}

			if product.Quantity != nil && *product.Quantity >= 0 {
				if err := services.UpdateProductQuantity(ctx, txProductRepo, product.ProductId, vendorId, quantitySum); err != nil {
					return err
				}
			}
		}

		return nil

	})

	if err != nil {
		return nil, err
	}

	return stockProductsRes, nil

}

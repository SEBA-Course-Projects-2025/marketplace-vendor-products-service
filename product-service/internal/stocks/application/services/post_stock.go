package services

import (
	"context"
	"dev-vendor/product-service/internal/products/application/services"
	productDomain "dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PostStock(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, db *gorm.DB, stockReq dtos.StockRequest, vendorId uuid.UUID) (dtos.PostStockResponse, error) {

	var stockResponse dtos.PostStockResponse

	err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txStockRepo := stockRepo.WithTx(tx)
		txProductRepo := productRepo.WithTx(tx)

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

			quantitySum, err := GetQuantitySum(ctx, txStockRepo, product.ProductId, vendorId)

			if err != nil {
				return err
			}

			if err := services.UpdateProductQuantity(ctx, txProductRepo, product.ProductId, vendorId, quantitySum); err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return dtos.PostStockResponse{}, err
	}

	return stockResponse, nil

}

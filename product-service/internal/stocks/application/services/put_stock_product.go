package services

import (
	"context"
	"dev-vendor/product-service/internal/products/application/services"
	productDomain "dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func PutStockProduct(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, stockProductReq dtos.PutStockProductRequest, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) error {

	if err := stockRepo.CheckProduct(ctx, productId, vendorId); err != nil {
		return err
	}

	if err := services.UpdateProductQuantity(ctx, productRepo, productId, vendorId, stockProductReq.Quantity); err != nil {
		return err
	}

	existingStock, err := stockRepo.FindById(ctx, stockId, vendorId)
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

	return stockRepo.UpdateStockProduct(ctx, updatedStockProduct)

}

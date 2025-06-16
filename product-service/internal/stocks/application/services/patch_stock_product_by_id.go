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

func PatchStockProductById(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, stockProductReq dtos.PatchStockProductRequest, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) (dtos.StockProductInfo, error) {

	if productId != uuid.Nil {

		if err := stockRepo.CheckProduct(ctx, productId, vendorId); err != nil {
			return dtos.StockProductInfo{}, err
		}

		if stockProductReq.Quantity != nil && *stockProductReq.Quantity >= 0 {
			if err := services.UpdateProductQuantity(ctx, productRepo, productId, vendorId, *stockProductReq.Quantity); err != nil {
				return dtos.StockProductInfo{}, err
			}
		}
	}

	existingStock, err := stockRepo.FindById(ctx, stockId, vendorId)
	if err != nil {
		return dtos.StockProductInfo{}, err
	}

	var existingStockProduct *models.StocksProduct
	for i := range existingStock.StocksProducts {
		if existingStock.StocksProducts[i].ProductId == productId {
			existingStockProduct = &existingStock.StocksProducts[i]
			break
		}
	}

	updatedStockProduct := dtos.ModifyStockProductWithDto(existingStockProduct, stockProductReq)

	updatedStockProduct, err = stockRepo.PatchStockProductId(ctx, updatedStockProduct)

	return dtos.StocksProductToStockProductInfo(updatedStockProduct), nil

}

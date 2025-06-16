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

func PatchStockProducts(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, stockProductReq []dtos.PatchStockManyProductsRequest, stockId uuid.UUID, vendorId uuid.UUID) ([]dtos.StockProductInfo, error) {

	existingStock, err := stockRepo.FindById(ctx, stockId, vendorId)
	if err != nil {
		return nil, err
	}

	var modifiedProducts []models.StocksProduct

	var stockProductsRes []dtos.StockProductInfo

	for _, product := range stockProductReq {

		if err := stockRepo.CheckProduct(ctx, product.ProductId, vendorId); err != nil {
			return nil, err
		}

		if product.Quantity != nil && *product.Quantity >= 0 {
			if err := services.UpdateProductQuantity(ctx, productRepo, product.ProductId, vendorId, *product.Quantity); err != nil {
				return nil, err
			}
		}

		for i := range existingStock.StocksProducts {
			if existingStock.StocksProducts[i].ProductId == product.ProductId {
				updatedStockProduct := dtos.ModifyStockManyProductsWithDto(&existingStock.StocksProducts[i], product)
				modifiedProducts = append(modifiedProducts, *updatedStockProduct)
				break
			}
		}

	}

	updatedProducts, err := stockRepo.PatchStockProducts(ctx, modifiedProducts)
	if err != nil {
		return nil, err
	}

	for _, updProduct := range updatedProducts {
		stockProductsRes = append(stockProductsRes, dtos.StocksProductToStockProductInfo(&updProduct))
	}

	return stockProductsRes, nil

}

package services

import (
	"context"
	"dev-vendor/product-service/internal/products/application/services"
	productDomain "dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
)

func PostStock(ctx context.Context, repo domain.StockRepository, productRepo productDomain.ProductRepository, stockReq dtos.StockRequest, vendorId uuid.UUID) (dtos.PostStockResponse, error) {

	if _, err := repo.CheckLocation(ctx, stockReq.LocationId); err != nil {
		return dtos.PostStockResponse{}, err
	}

	for _, product := range stockReq.Products {

		if err := repo.CheckProduct(ctx, product.ProductId, vendorId); err != nil {
			return dtos.PostStockResponse{}, err
		}

		if err := services.UpdateProductQuantity(ctx, productRepo, product.ProductId, vendorId, product.Quantity); err != nil {
			return dtos.PostStockResponse{}, err
		}
	}

	newStock, err := dtos.PostStockRequestToStock(stockReq, vendorId)
	if err != nil {
		return dtos.PostStockResponse{}, err
	}

	createdStock, err := repo.Create(ctx, newStock, vendorId)
	if err != nil {
		return dtos.PostStockResponse{}, err
	}

	return dtos.PostStockToStockResponse(createdStock), nil

}

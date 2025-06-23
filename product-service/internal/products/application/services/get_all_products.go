package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"github.com/google/uuid"
)

func GetAllProducts(ctx context.Context, repo domain.ProductRepository, params dtos.ProductQueryParams, vendorId uuid.UUID) ([]dtos.GetProductsResponse, error) {

	products, err := repo.FindAll(ctx, params, vendorId)

	if err != nil {
		return nil, err
	}

	return dtos.ProductsToDto(products), nil

}

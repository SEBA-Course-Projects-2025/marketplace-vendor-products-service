package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"github.com/google/uuid"
)

func GetProductById(ctx context.Context, repo domain.ProductRepository, id uuid.UUID, vendorId uuid.UUID) (dtos.OneProductResponse, error) {

	product, err := repo.FindById(ctx, id, vendorId)

	if err != nil {
		return dtos.OneProductResponse{}, err
	}

	return dtos.ProductToDto(*product), nil

}

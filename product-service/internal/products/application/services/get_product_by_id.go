package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/google/uuid"
)

func GetProductById(ctx context.Context, repo domain.ProductRepository, id uuid.UUID) (dtos.OneProductResponse, error) {

	ctx, span := tracer.Tracer.Start(ctx, "GetOneProduct")
	defer span.End()

	product, err := repo.FindById(ctx, id)

	if err != nil {
		return dtos.OneProductResponse{}, err
	}

	return dtos.ProductToDto(product), nil

}

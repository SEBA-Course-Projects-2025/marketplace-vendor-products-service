package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/google/uuid"
)

func GetProductBySlug(ctx context.Context, repo domain.ProductRepository, slug string, vendorId uuid.UUID) (dtos.OneProductResponse, error) {

	ctx, span := tracer.Tracer.Start(ctx, "GetOneProductBySlug")
	defer span.End()

	product, err := repo.FindBySlug(ctx, slug, vendorId)

	if err != nil {
		return dtos.OneProductResponse{}, err
	}

	return dtos.ProductToDto(product), nil

}

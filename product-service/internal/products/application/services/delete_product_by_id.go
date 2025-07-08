package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/google/uuid"
)

func DeleteProductById(ctx context.Context, repo domain.ProductRepository, id uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteOneProduct")
	defer span.End()

	return repo.Transaction(func(txRepo domain.ProductRepository) error {
		return txRepo.DeleteById(ctx, id, vendorId)
	})

}

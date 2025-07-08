package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/google/uuid"
)

func DeleteManyProducts(ctx context.Context, repo domain.ProductRepository, ids []uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteAllProducts")
	defer span.End()

	return repo.Transaction(func(txRepo domain.ProductRepository) error {
		return txRepo.DeleteMany(ctx, ids, vendorId)
	})

}

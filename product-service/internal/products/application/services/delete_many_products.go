package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"github.com/google/uuid"
)

func DeleteManyProducts(ctx context.Context, repo domain.ProductRepository, ids []uuid.UUID, vendorId uuid.UUID) error {
	return repo.Transaction(func(txRepo domain.ProductRepository) error {
		return txRepo.DeleteMany(ctx, ids, vendorId)
	})
}

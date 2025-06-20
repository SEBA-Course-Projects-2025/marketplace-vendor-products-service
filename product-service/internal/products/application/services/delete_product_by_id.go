package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"github.com/google/uuid"
)

func DeleteProductById(ctx context.Context, repo domain.ProductRepository, id uuid.UUID, vendorId uuid.UUID) error {
	return repo.Transaction(func(txRepo domain.ProductRepository) error {
		return txRepo.DeleteById(ctx, id, vendorId)
	})

}

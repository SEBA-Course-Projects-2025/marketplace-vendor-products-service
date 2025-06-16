package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"github.com/google/uuid"
)

func UpdateProductQuantity(ctx context.Context, repo domain.ProductRepository, productId uuid.UUID, vendorId uuid.UUID, quantity int) error {

	product, err := repo.FindById(ctx, productId, vendorId)

	if err != nil {
		return err
	}

	product.Quantity += quantity - product.Quantity

	return repo.Update(ctx, product)

}

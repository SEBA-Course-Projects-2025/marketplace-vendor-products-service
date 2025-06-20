package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"github.com/google/uuid"
)

func PutProduct(ctx context.Context, repo domain.ProductRepository, id uuid.UUID, productReq dtos.ProductRequest, vendorId uuid.UUID) error {

	return repo.Transaction(func(txRepo domain.ProductRepository) error {
		existingProduct, err := txRepo.FindById(ctx, id, vendorId)

		if err != nil {
			return err
		}

		existingProduct = dtos.UpdateProductWithDto(existingProduct, productReq)
		existingProduct.VendorId = vendorId

		return txRepo.Update(ctx, existingProduct)
	})

}

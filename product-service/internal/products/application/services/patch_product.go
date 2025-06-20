package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"github.com/google/uuid"
)

func PatchProduct(ctx context.Context, repo domain.ProductRepository, id uuid.UUID, productReq dtos.ProductPatchRequest, vendorId uuid.UUID) (dtos.OneProductResponse, error) {

	var productResponse dtos.OneProductResponse

	err := repo.Transaction(func(txRepo domain.ProductRepository) error {
		existingProduct, err := txRepo.FindById(ctx, id, vendorId)

		if err != nil {
			return err
		}

		existingProduct = dtos.PatchDtoToProduct(existingProduct, productReq)

		existingProduct.VendorId = vendorId

		existingProduct, err = txRepo.Patch(ctx, existingProduct)

		if err != nil {
			return err
		}

		productResponse = dtos.ProductToDto(*existingProduct)
		return nil
	})

	if err != nil {
		return dtos.OneProductResponse{}, err
	}

	return productResponse, nil
}

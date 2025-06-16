package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/utils"
	"github.com/google/uuid"
)

func PatchProduct(ctx context.Context, repo domain.ProductRepository, id uuid.UUID, productReq dtos.ProductPatchRequest, vendorId uuid.UUID) (dtos.OneProductResponse, error) {

	existingProduct, err := repo.FindById(ctx, id, vendorId)

	if err != nil {
		return dtos.OneProductResponse{}, err
	}

	existingProduct = dtos.PatchDtoToProduct(existingProduct, productReq)

	existingProduct.VendorId = vendorId

	existingProduct, err = repo.Patch(ctx, existingProduct)

	if err != nil {
		return dtos.OneProductResponse{}, utils.ErrorHandler(err, "Error updating product")
	}

	return dtos.ProductToDto(*existingProduct), nil
}

package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"github.com/google/uuid"
)

func PostProduct(ctx context.Context, repo domain.ProductRepository, productReq dtos.ProductRequest, vendorId uuid.UUID) (dtos.OneProductResponse, error) {

	var productResponse dtos.OneProductResponse

	err := repo.Transaction(func(txRepo domain.ProductRepository) error {

		newProduct := dtos.PostDtoToProduct(productReq, vendorId)

		product, err := txRepo.Create(ctx, &newProduct, vendorId)

		if err != nil {
			return err
		}

		productResponse = dtos.ProductToDto(*product)

		return nil
	})

	if err != nil {
		return dtos.OneProductResponse{}, nil
	}

	return productResponse, nil

}

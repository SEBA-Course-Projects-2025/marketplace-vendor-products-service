package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"github.com/google/uuid"
)

func PostProduct(ctx context.Context, repo domain.ProductRepository, productReq dtos.ProductRequest, vendorId uuid.UUID) (dtos.OneProductResponse, error) {

	newProduct := dtos.PostDtoToProduct(productReq, vendorId)

	product, err := repo.Create(ctx, &newProduct, vendorId)

	if err != nil {
		return dtos.OneProductResponse{}, err
	}

	productResponse := dtos.ProductToDto(*product)

	return productResponse, nil

}

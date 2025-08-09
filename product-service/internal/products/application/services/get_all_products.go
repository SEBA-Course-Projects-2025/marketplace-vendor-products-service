package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetAllProducts(ctx context.Context, repo domain.ProductRepository, params dtos.ProductQueryParams, vendorId uuid.UUID) ([]dtos.GetProductsResponse, error) {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Starting GetAllProducts application service")

	ctx, span := tracer.Tracer.Start(ctx, "GetAllProducts")
	defer span.End()

	products, err := repo.FindAll(ctx, params, vendorId)

	if err != nil {
		return nil, err
	}

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Successfully get paginated list of products")
	return dtos.ProductsToDto(products), nil

}

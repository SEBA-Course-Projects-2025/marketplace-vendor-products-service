package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetProductBySlug(ctx context.Context, repo domain.ProductRepository, slug string, vendorId uuid.UUID) (dtos.OneProductResponse, error) {

	logrus.WithFields(logrus.Fields{
		"slug":     slug,
	}).Info("Starting GetProductBySlug application service")

	ctx, span := tracer.Tracer.Start(ctx, "GetOneProductBySlug")
	defer span.End()

	product, err := repo.FindBySlug(ctx, slug, vendorId)

	if err != nil {
		return dtos.OneProductResponse{}, err
	}

	logrus.WithFields(logrus.Fields{
		"slug":     slug,
	}).Info("Successfully get product by slug")
	return dtos.ProductToDto(product), nil

}

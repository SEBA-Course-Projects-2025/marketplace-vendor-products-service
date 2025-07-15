package services

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GetProductById(ctx context.Context, repo domain.ProductRepository, id uuid.UUID) (dtos.OneProductResponse, error) {

	logrus.WithFields(logrus.Fields{
		"productId":     id,
	}).Info("Starting GetProductById application service")

	ctx, span := tracer.Tracer.Start(ctx, "GetOneProductById")
	defer span.End()

	product, err := repo.FindById(ctx, id)

	if err != nil {
		return dtos.OneProductResponse{}, err
	}

	logrus.WithFields(logrus.Fields{
		"productId":     id,
	}).Info("Successfully get product by id")
	return dtos.ProductToDto(product), nil

}

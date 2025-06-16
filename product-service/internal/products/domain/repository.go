package domain

import (
	"context"
	"dev-vendor/product-service/internal/products/domain/models"
	"dev-vendor/product-service/internal/products/dtos"
	"github.com/google/uuid"
)

type ProductRepository interface {
	FindById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) (*models.Product, error)
	FindAll(ctx context.Context, params dtos.ProductQueryParams, vendorId uuid.UUID) (*[]models.Product, error)
	Create(ctx context.Context, newProduct *models.Product, vendorId uuid.UUID) (*models.Product, error)
	Update(ctx context.Context, updatedProduct *models.Product) error
	Patch(ctx context.Context, modifiedProduct *models.Product) (*models.Product, error)
	DeleteById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) error
	DeleteMany(ctx context.Context, ids []uuid.UUID, vendorId uuid.UUID) error
}

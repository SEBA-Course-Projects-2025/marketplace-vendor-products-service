package domain

import (
	"context"
	"dev-vendor/product-service/internal/products/domain/productModels"
	"dev-vendor/product-service/internal/products/dtos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindById(ctx context.Context, id uuid.UUID) (*productModels.Product, error)
	FindBySlug(ctx context.Context, slug string, vendorId uuid.UUID) (*productModels.Product, error)
	FindAll(ctx context.Context, params dtos.ProductQueryParams, vendorId uuid.UUID) ([]productModels.Product, error)
	FindAllTags(ctx context.Context) ([]productModels.Tag, error)
	Create(ctx context.Context, newProduct *productModels.Product, vendorId uuid.UUID) (*productModels.Product, error)
	Update(ctx context.Context, updatedProduct *productModels.Product) error
	Patch(ctx context.Context, modifiedProduct *productModels.Product) (*productModels.Product, error)
	DeleteById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) error
	DeleteBySlug(ctx context.Context, slug string, vendorId uuid.UUID) error
	DeleteProductImages(ctx context.Context, product *productModels.Product) error
	DeleteProductTags(ctx context.Context, product *productModels.Product) error
	DeleteMany(ctx context.Context, ids []uuid.UUID, vendorId uuid.UUID) error
	Transaction(fn func(txRepo ProductRepository) error) error
	WithTx(tx *gorm.DB) ProductRepository
}

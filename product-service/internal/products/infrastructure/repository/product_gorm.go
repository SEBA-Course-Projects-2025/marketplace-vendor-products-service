package repository

import (
	"context"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/domain/productModels"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type GormProductRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{db: db}
}

func (gpr *GormProductRepository) FindById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) (*productModels.Product, error) {

	var product productModels.Product

	if err := gpr.db.WithContext(ctx).Preload("Images").Preload("Tags").First(&product, "id = ? AND vendor_id = ?", id, vendorId).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error getting product data")
	}

	return &product, nil

}

func (gpr *GormProductRepository) FindAll(ctx context.Context, params dtos.ProductQueryParams, vendorId uuid.UUID) ([]productModels.Product, error) {

	var products []productModels.Product

	db := gpr.db.WithContext(ctx)

	db = db.Where("vendor_id = ?", vendorId)

	db = db.Preload("Images").Preload("Tags")

	if params.Category != "" {
		db = db.Where("category = ?", params.Category)
	}

	if params.MinPrice > 0 {
		db = db.Where("price >= ?", params.MinPrice)
	}

	if params.MaxPrice > 0 {
		db = db.Where("price <= ?", params.MaxPrice)
	}

	if params.Search != "" {
		db = db.Where("name ILIKE ?", "%"+params.Search+"%")
	}

	allowedSortBy := map[string]bool{
		"price":    true,
		"name":     true,
		"quantity": true,
	}

	orderField := "name"

	if allowedSortBy[params.SortBy] {
		orderField = params.SortBy
	}

	orderDir := "asc"

	if params.SortOrder == "desc" {
		orderDir = "desc"
	}

	db = db.Order(orderField + " " + orderDir)

	if err := db.Limit(params.Limit).Offset(params.Offset).Find(&products).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error getting paginated products data")
	}

	return products, nil

}

func (gpr *GormProductRepository) Create(ctx context.Context, newProduct *productModels.Product, vendorId uuid.UUID) (*productModels.Product, error) {

	newProduct.VendorId = vendorId

	if err := gpr.db.WithContext(ctx).Create(newProduct).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error creating product")
	}

	return newProduct, nil

}

func (gpr *GormProductRepository) Update(ctx context.Context, updatedProduct *productModels.Product) error {

	res := gpr.db.WithContext(ctx).Save(updatedProduct)

	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error updating product")
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gpr *GormProductRepository) Patch(ctx context.Context, modifiedProduct *productModels.Product) (*productModels.Product, error) {

	res := gpr.db.WithContext(ctx).Save(modifiedProduct)

	if res.Error != nil {
		return nil, utils.ErrorHandler(res.Error, "Error updating product")
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return modifiedProduct, nil

}

func (gpr *GormProductRepository) DeleteById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) error {

	res := gpr.db.WithContext(ctx).Where("id = ? AND vendor_id = ?", id, vendorId).Delete(&productModels.Product{})
	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error deleting product")
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gpr *GormProductRepository) DeleteMany(ctx context.Context, ids []uuid.UUID, vendorId uuid.UUID) error {

	res := gpr.db.WithContext(ctx).Where("vendor_id = ? AND id IN ?", vendorId, ids).Delete(&productModels.Product{})

	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error deleting product")
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gpr *GormProductRepository) WithTx(tx *gorm.DB) domain.ProductRepository {
	return &GormProductRepository{
		db: tx,
	}
}

func (gpr *GormProductRepository) Transaction(fn func(txRepo domain.ProductRepository) error) error {

	tx := gpr.db.Begin()
	if tx.Error != nil {
		log.Printf("Transaction begin error: %v", tx.Error)
		return tx.Error
	}

	repo := gpr.WithTx(tx)

	if err := fn(repo); err != nil {
		log.Printf("Transaction function error: %v", err)
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		log.Printf("Transaction commit error: %v", err)
		return err
	}

	return nil

}

package repository

import (
	"context"
	productsModels "dev-vendor/product-service/internal/products/domain/productModels"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/shared/utils"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormStockRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *GormStockRepository {
	return &GormStockRepository{db: db}
}

func (gsr *GormStockRepository) FindById(ctx context.Context, id uuid.UUID) (*models.Stock, error) {

	ctx, span := tracer.Tracer.Start(ctx, "FindById")
	defer span.End()

	var stock models.Stock

	if err := gsr.db.WithContext(ctx).Preload("Location").Preload("StocksProducts.Product.Images").First(&stock, "id = ?", id).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error getting stock data")
	}

	return &stock, nil
}

func (gsr *GormStockRepository) FindAll(ctx context.Context, params dtos.StockQueryParams, vendorId uuid.UUID) ([]models.Stock, error) {

	ctx, span := tracer.Tracer.Start(ctx, "FindAll")
	defer span.End()

	var stocks []models.Stock

	db := gsr.db.WithContext(ctx)

	db = db.Where("vendor_id = ?", vendorId)

	if params.LocationId != "" {
		db = db.Where("location_id = ?", params.LocationId)
	}

	orderField := "date_supplied"
	orderDir := "asc"

	if params.SortOrder == "desc" {
		orderDir = "desc"
	}

	db = db.Order(orderField + " " + orderDir)

	if err := db.Limit(params.Limit).Offset(params.Offset).Find(&stocks).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error getting paginated stocks data")
	}

	return stocks, nil

}

func (gsr *GormStockRepository) Create(ctx context.Context, newStock *models.Stock, vendorId uuid.UUID) (*models.Stock, error) {

	ctx, span := tracer.Tracer.Start(ctx, "Create")
	defer span.End()

	newStock.VendorId = vendorId

	if err := gsr.db.WithContext(ctx).Create(newStock).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error creating stock")
	}

	return newStock, nil

}

func (gsr *GormStockRepository) UpdateStock(ctx context.Context, updatedStock *models.Stock) error {

	ctx, span := tracer.Tracer.Start(ctx, "UpdateStock")
	defer span.End()

	res := gsr.db.WithContext(ctx).Save(updatedStock)

	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error updating stock")
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gsr *GormStockRepository) UpdateStockProduct(ctx context.Context, updatedStockProduct *models.StocksProduct) error {

	ctx, span := tracer.Tracer.Start(ctx, "UpdateStockProduct")
	defer span.End()

	res := gsr.db.WithContext(ctx).Save(updatedStockProduct)

	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error updating stock product")
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gsr *GormStockRepository) PatchStockId(ctx context.Context, modifiedStock *models.Stock) (*models.Stock, error) {

	ctx, span := tracer.Tracer.Start(ctx, "PatchStockId")
	defer span.End()

	res := gsr.db.WithContext(ctx).Save(modifiedStock)

	if res.Error != nil {
		return nil, utils.ErrorHandler(res.Error, "Error modifying stock")
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return modifiedStock, nil
}

func (gsr *GormStockRepository) PatchStockProductId(ctx context.Context, modifiedStockProduct *models.StocksProduct) (*models.StocksProduct, error) {

	ctx, span := tracer.Tracer.Start(ctx, "PatchStockProductId")
	defer span.End()

	res := gsr.db.WithContext(ctx).Save(modifiedStockProduct)

	if res.Error != nil {
		return nil, utils.ErrorHandler(res.Error, "Error updating stock product")
	}

	if res.RowsAffected == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return modifiedStockProduct, nil

}

func (gsr *GormStockRepository) PatchStockProducts(ctx context.Context, modifiedStockProducts []models.StocksProduct) ([]models.StocksProduct, error) {

	ctx, span := tracer.Tracer.Start(ctx, "PatchStockProducts")
	defer span.End()

	for i := range modifiedStockProducts {

		stockProduct := &modifiedStockProducts[i]

		res := gsr.db.WithContext(ctx).Save(stockProduct)
		if res.Error != nil {
			return nil, utils.ErrorHandler(res.Error, "Error updating stock products")
		}
		if res.RowsAffected == 0 {
			return nil, gorm.ErrRecordNotFound
		}
	}

	return modifiedStockProducts, nil

}

func (gsr *GormStockRepository) DeleteStockById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteStockById")
	defer span.End()

	res := gsr.db.WithContext(ctx).Where("id = ? AND vendor_id = ?", id, vendorId).Delete(&models.Stock{})
	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error deleting stock")
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gsr *GormStockRepository) DeleteStockProductById(ctx context.Context, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteStockProductById")
	defer span.End()

	var stock models.Stock

	if err := gsr.db.WithContext(ctx).First(&stock, "id = ? AND vendor_id = ?", stockId, vendorId).Error; err != nil {
		return utils.ErrorHandler(err, "Invalid stock data")
	}

	res := gsr.db.WithContext(ctx).Where("stock_id = ? AND product_id = ?", stockId, productId).Delete(&models.StocksProduct{})
	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error deleting stock product")
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gsr *GormStockRepository) DeleteManyStocks(ctx context.Context, ids []uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteManyStocks")
	defer span.End()

	res := gsr.db.WithContext(ctx).Where("vendor_id = ? AND id IN ?", vendorId, ids).Delete(&models.Stock{})

	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error deleting stock")
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gsr *GormStockRepository) DeleteManyStockProducts(ctx context.Context, ids []uuid.UUID, stockId uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteManyStockProducts")
	defer span.End()

	var stock models.Stock
	if err := gsr.db.WithContext(ctx).First(&stock, "id = ? AND vendor_id = ?", stockId, vendorId).Error; err != nil {
		return utils.ErrorHandler(err, "Invalid stock data")
	}

	res := gsr.db.WithContext(ctx).Where("stock_id = ? AND product_id IN ?", stockId, ids).Delete(&models.StocksProduct{})

	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error deleting stock products")
	}

	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gsr *GormStockRepository) CheckProduct(ctx context.Context, productId uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "CheckProduct")
	defer span.End()

	var product productsModels.Product

	if err := gsr.db.WithContext(ctx).First(&product, "id = ? AND vendor_id = ?", productId, vendorId).Error; err != nil {
		return utils.ErrorHandler(err, "Error: no such product")
	}

	return nil

}

func (gsr *GormStockRepository) CheckLocation(ctx context.Context, locationId uuid.UUID) (*models.StocksLocation, error) {

	ctx, span := tracer.Tracer.Start(ctx, "CheckLocation")
	defer span.End()

	var location models.StocksLocation

	if err := gsr.db.WithContext(ctx).First(&location, "id = ?", locationId).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error: no such location")
	}

	return &location, nil

}

func (gsr *GormStockRepository) FindProductStocksQuantities(ctx context.Context, productId uuid.UUID) ([]models.StocksProduct, error) {

	ctx, span := tracer.Tracer.Start(ctx, "FindProductStocksQuantities")
	defer span.End()

	var productStocks []models.StocksProduct

	if err := gsr.db.WithContext(ctx).Where("product_id = ?", productId).Preload("Stock").Find(&productStocks).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error: no such products in stocks")
	}

	return productStocks, nil

}

func (gsr *GormStockRepository) FindAllStockProducts(ctx context.Context, params dtos.StockProductsQueryParams, vendorId uuid.UUID) ([]models.StocksProduct, error) {

	ctx, span := tracer.Tracer.Start(ctx, "FindAllStockProducts")
	defer span.End()

	var stockProducts []models.StocksProduct

	db := gsr.db.WithContext(ctx)

	db = db.Preload("Stock", "vendor_id = ?", vendorId).Preload("Product").Preload("Product").Preload("Product.Images")

	allowedSortBy := map[string]string{
		"unit-cost": "unit_cost",
		"quantity":  "quantity",
	}

	orderField := "unit_cost"

	if field, ok := allowedSortBy[params.SortBy]; ok {
		orderField = field
	}

	orderDir := "asc"

	if params.SortOrder == "desc" {
		orderDir = "desc"
	}

	db = db.Order(orderField + " " + orderDir)

	if err := db.Limit(params.Limit).Offset(params.Offset).Find(&stockProducts).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error getting paginated stock products data")
	}

	return stockProducts, nil

}

func (gsr *GormStockRepository) WithTx(tx *gorm.DB) domain.StockRepository {
	return &GormStockRepository{
		db: tx,
	}
}

func (gsr *GormStockRepository) Transaction(fn func(txRepo domain.StockRepository) error) error {

	tx := gsr.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	repo := gsr.WithTx(tx)

	if err := fn(repo); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil

}

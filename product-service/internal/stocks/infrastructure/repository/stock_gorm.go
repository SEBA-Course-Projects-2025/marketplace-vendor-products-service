package repository

import (
	"context"
	productsModels "dev-vendor/product-service/internal/products/domain/models"
	"dev-vendor/product-service/internal/shared/utils"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormStockRepository struct {
	Db *gorm.DB
}

func (gsr *GormStockRepository) FindById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) (*models.Stock, error) {

	var stock models.Stock

	if err := gsr.Db.WithContext(ctx).Preload("Location").Preload("StocksProducts.Product.Images").First(&stock, "id = ? AND vendor_id = ?", id, vendorId).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error getting stock data")
	}

	return &stock, nil
}

func (gsr *GormStockRepository) FindAll(ctx context.Context, params dtos.StockQueryParams, vendorId uuid.UUID) (*[]models.Stock, error) {

	var stocks *[]models.Stock

	db := gsr.Db.WithContext(ctx)

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

	newStock.VendorId = vendorId

	if err := gsr.Db.WithContext(ctx).Create(newStock).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error creating stock")
	}

	return newStock, nil

}

func (gsr *GormStockRepository) UpdateStock(ctx context.Context, updatedStock *models.Stock) error {

	if err := gsr.Db.WithContext(ctx).Save(updatedStock).Error; err != nil {
		return utils.ErrorHandler(err, "Error updating stock")
	}

	return nil

}

func (gsr *GormStockRepository) UpdateStockProduct(ctx context.Context, updatedStockProduct *models.StocksProduct) error {

	if err := gsr.Db.WithContext(ctx).Save(updatedStockProduct).Error; err != nil {
		return utils.ErrorHandler(err, "Error updating stock product")
	}

	return nil

}

func (gsr *GormStockRepository) PatchStockId(ctx context.Context, modifiedStock *models.Stock) (*models.Stock, error) {

	if err := gsr.Db.WithContext(ctx).Save(modifiedStock).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error modifying stock")
	}

	return modifiedStock, nil
}

func (gsr *GormStockRepository) PatchStockProductId(ctx context.Context, modifiedStockProduct *models.StocksProduct) (*models.StocksProduct, error) {

	if err := gsr.Db.WithContext(ctx).Save(modifiedStockProduct).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error updating stock product")
	}

	return modifiedStockProduct, nil

}

func (gsr *GormStockRepository) PatchStockProducts(ctx context.Context, modifiedStockProducts []models.StocksProduct) ([]models.StocksProduct, error) {

	err := gsr.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		for i := range modifiedStockProducts {

			stockProduct := &modifiedStockProducts[i]

			if err := tx.Save(stockProduct).Error; err != nil {
				return utils.ErrorHandler(err, "Error updating stock products")
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return modifiedStockProducts, nil

}

func (gsr *GormStockRepository) DeleteStockById(ctx context.Context, id uuid.UUID, vendorId uuid.UUID) error {

	res := gsr.Db.WithContext(ctx).Where("id = ? AND vendor_id = ?", id, vendorId).Delete(&models.Stock{})
	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error deleting stock")
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gsr *GormStockRepository) DeleteStockProductById(ctx context.Context, stockId uuid.UUID, productId uuid.UUID, vendorId uuid.UUID) error {

	var stock models.Stock

	if err := gsr.Db.WithContext(ctx).First(&stock, "id = ? AND vendor_id = ?", stockId, vendorId).Error; err != nil {
		return utils.ErrorHandler(err, "Invalid stock data")
	}

	res := gsr.Db.WithContext(ctx).Where("stock_id = ? AND product_id = ?", stockId, productId).Delete(&models.StocksProduct{})
	if res.Error != nil {
		return utils.ErrorHandler(res.Error, "Error deleting stock product")
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil

}

func (gsr *GormStockRepository) DeleteManyStocks(ctx context.Context, ids []uuid.UUID, vendorId uuid.UUID) error {

	return gsr.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		res := tx.Where("vendor_id = ? AND id IN ?", vendorId, ids).Delete(&models.Stock{})

		if res.Error != nil {
			return utils.ErrorHandler(res.Error, "Error deleting stock")
		}

		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil

	})

}

func (gsr *GormStockRepository) DeleteManyStockProducts(ctx context.Context, ids []uuid.UUID, stockId uuid.UUID, vendorId uuid.UUID) error {

	return gsr.Db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var stock models.Stock
		if err := tx.First(&stock, "id = ? AND vendor_id = ?", stockId, vendorId).Error; err != nil {
			return utils.ErrorHandler(err, "Invalid stock data")
		}

		res := tx.Where("stock_id = ? AND product_id IN ?", stockId, ids).Delete(&models.StocksProduct{})

		if res.Error != nil {
			return utils.ErrorHandler(res.Error, "Error deleting stock products")
		}

		if res.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil

	})

}

func (gsr *GormStockRepository) CheckProduct(ctx context.Context, productId uuid.UUID, vendorId uuid.UUID) error {

	var product productsModels.Product

	if err := gsr.Db.WithContext(ctx).First(&product, "id = ? AND vendor_id = ?", productId, vendorId).Error; err != nil {
		return utils.ErrorHandler(err, "Error: no such product")
	}

	return nil

}

func (gsr *GormStockRepository) CheckLocation(ctx context.Context, locationId uuid.UUID) (*models.StocksLocation, error) {

	var location models.StocksLocation

	if err := gsr.Db.WithContext(ctx).First(&location, "id = ?", locationId).Error; err != nil {
		return nil, utils.ErrorHandler(err, "Error: no such location")
	}

	return &location, nil

}

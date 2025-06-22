package models

import (
	"dev-vendor/product-service/internal/products/domain/productModels"
	"github.com/google/uuid"
	"time"
)

type StocksLocation struct {
	Id      uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	City    string    `json:"city" gorm:"column:city;type:varchar(255);not null"`
	Address string    `json:"address" gorm:"column:address;type:varchar(255);not null"`
	Stocks  []Stock   `json:"stocks" gorm:"foreignKey:LocationId;references:Id"`
}

type Stock struct {
	Id             uuid.UUID       `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	VendorId       uuid.UUID       `json:"vendor_id" gorm:"column:vendor_id;type:uuid"`
	DateSupplied   time.Time       `json:"date_supplied" gorm:"column:date_supplied;type:date;not null"`
	LocationId     uuid.UUID       `json:"location_id" gorm:"column:location_id;type:uuid"`
	Location       StocksLocation  `json:"location" gorm:"foreignKey:LocationId;references:Id"`
	StocksProducts []StocksProduct `json:"stocks_products" gorm:"foreignKey:StockId;references:Id"`
	CreatedAt      time.Time       `gorm:"column:created_at;type:timestamp"`
	UpdatedAt      time.Time       `gorm:"column:updated_at;type:timestamp"`
}

type StocksProduct struct {
	StockId   uuid.UUID             `json:"stock_id" gorm:"column:stock_id;type:uuid;primaryKey"`
	ProductId uuid.UUID             `json:"product_id" gorm:"column:product_id;type:uuid;primaryKey"`
	Quantity  int                   `json:"quantity" gorm:"column:quantity;type:int;not null"`
	UnitCost  float64               `json:"unit_cost" gorm:"column:unit_cost;type:numeric(12, 2);not null"`
	CreatedAt time.Time             `gorm:"column:created_at;type:timestamp"`
	UpdatedAt time.Time             `gorm:"column:updated_at;type:timestamp"`
	Stock     Stock                 `json:"stock" gorm:"foreignKey:StockId;references:Id"`
	Product   productModels.Product `json:"product" gorm:"foreignKey:ProductId;references:Id"`
}

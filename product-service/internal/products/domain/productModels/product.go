package productModels

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	Id          uuid.UUID       `json:"id" gorm:"column:id;type:uuid;primaryKey"`
	VendorId    uuid.UUID       `json:"vendor_id" gorm:"column:vendor_id;type:uuid;not null"`
	Name        string          `json:"name" gorm:"column:name;type:varchar(255);not null"`
	Description string          `json:"description" gorm:"column:description;type:text;not null"`
	Price       float64         `json:"price" gorm:"column:price;type:numeric(12, 2);not null"`
	Category    string          `json:"category" gorm:"column:category;type:varchar(255);not null"`
	Slug        string          `json:"slug" gorm:"column:slug;type:varchar(255);not null"`
	Images      []ProductsImage `json:"images" gorm:"foreignKey:ProductId"`
	Tags        []Tag           `json:"tags" gorm:"many2many:products_tags;"`
	Quantity    int             `json:"quantity" gorm:"column:quantity;type:int;not null"`
	CreatedAt   time.Time       `gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time       `gorm:"column:updated_at;type:timestamp"`
	DeletedAt   *time.Time      `gorm:"column:deleted_at;type:timestamp"`
}

type ProductsImage struct {
	Id        uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey;not null"`
	ImageUrl  string    `json:"image_url" gorm:"column:image_url;type:text;not null"`
	ProductId uuid.UUID `json:"product_id" gorm:"column:product_id;type:uuid;not null"`
}

type Tag struct {
	Id      uuid.UUID `json:"id" gorm:"column:id;type:uuid;primaryKey;not null"`
	TagName string    `json:"tag_name" gorm:"column:tag_name;type:text;not null"`
}

type ProductsTag struct {
	ProductId uuid.UUID `gorm:"column:product_id;type:uuid;primaryKey"`
	TagId     uuid.UUID `gorm:"column:tag_id;type:uuid;primaryKey"`
}

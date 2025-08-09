package dtos

import (
	"dev-vendor/product-service/internal/event/domain/models"
	"dev-vendor/product-service/internal/products/domain/productModels"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type OneProductResponse struct {
	Id          uuid.UUID          `json:"id"`
	VendorId    uuid.UUID          `json:"vendor_id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Category    string             `json:"category"`
	Images      []ProductsImageDto `json:"images"`
	Tags        []TagDto           `json:"tags"`
	Quantity    int                `json:"quantity"`
	CreatedAt   time.Time          `json:"created_date"`
}

func ProductToDto(product *productModels.Product) OneProductResponse {
	return OneProductResponse{
		Id:          product.Id,
		VendorId:    product.VendorId,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Images:      ProductImagesToDto(product),
		Tags:        ProductTagsToDto(product),
		Quantity:    product.Quantity,
		CreatedAt:   product.CreatedAt,
	}
}

type GetProductsResponse struct {
	Id        uuid.UUID        `json:"id"`
	VendorId  uuid.UUID        `json:"vendor_id"`
	Name      string           `json:"name"`
	Price     float64          `json:"price"`
	Category  string           `json:"category"`
	Image     ProductsImageDto `json:"image"`
	Quantity  int              `json:"quantity"`
	CreatedAt time.Time        `json:"created_date"`
}

func ProductsToDto(products []productModels.Product) []GetProductsResponse {

	var productsResponse []GetProductsResponse

	for _, product := range products {

		var image productModels.ProductsImage

		if len(product.Images) > 0 {
			image = product.Images[0]
		} else {
			image = productModels.ProductsImage{}
		}

		imageDto := ProductsImageDto{
			Id:        image.Id,
			ImageUrl:  image.ImageUrl,
			ProductId: image.ProductId,
		}

		productResponse := GetProductsResponse{
			Id:        product.Id,
			VendorId:  product.VendorId,
			Name:      product.Name,
			Price:     product.Price,
			Category:  product.Category,
			Image:     imageDto,
			Quantity:  product.Quantity,
			CreatedAt: product.CreatedAt,
		}
		productsResponse = append(productsResponse, productResponse)
	}

	return productsResponse
}

type ProductRequest struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Category    string   `json:"category"`
	Images      []string `json:"images"`
	Tags        []string `json:"tags"`
}

func ValidateNewProductReq(productReq ProductRequest) error {

	if productReq.Name == "" {
		logrus.Error("Product name is missing when creating a new one")
		return errors.New("product name is missing")
	}

	if productReq.Description == "" {
		logrus.Error("Product description is missing when creating a new one")
		return errors.New("product description is missing")
	}

	if productReq.Price <= 0 {
		logrus.Error("Product price is missing or invalid when creating a new one")
		return errors.New("product price is missing or invalid")
	}

	if productReq.Category == "" {
		logrus.Error("Product category is missing when creating a new one")
		return errors.New("product category is missing")
	}

	if len(productReq.Images) == 0 {
		logrus.Error("Product images are missing when creating a new one")
		return errors.New("product images are missing")
	}

	if len(productReq.Tags) == 0 {
		logrus.Error("Product tags are missing when creating a new one")
		return errors.New("product tags are missing")
	}

	return nil

}

func PostDtoToProduct(productReq ProductRequest, vendorId uuid.UUID) (productModels.Product, error) {

	if err := ValidateNewProductReq(productReq); err != nil {
		return productModels.Product{}, err
	}

	productId := uuid.New()

	images := make([]productModels.ProductsImage, 0, len(productReq.Images))

	for _, imageUrl := range productReq.Images {
		images = append(images, productModels.ProductsImage{
			Id:        uuid.New(),
			ImageUrl:  imageUrl,
			ProductId: productId,
		})
	}

	tags := make([]productModels.Tag, 0, len(productReq.Tags))

	for _, tag := range productReq.Tags {
		tags = append(tags, productModels.Tag{
			Id:      uuid.New(),
			TagName: tag,
		})
	}

	product := productModels.Product{
		Id:          productId,
		VendorId:    vendorId,
		Name:        productReq.Name,
		Description: productReq.Description,
		Price:       productReq.Price,
		Category:    productReq.Category,
		Quantity:    0,
		Images:      images,
		Tags:        tags,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return product, nil
}

func UpdateProductWithDto(existingProduct *productModels.Product, productReq ProductRequest, tags []productModels.Tag) *productModels.Product {

	images := make([]productModels.ProductsImage, 0, len(productReq.Images))

	for _, imageUrl := range productReq.Images {
		images = append(images, productModels.ProductsImage{
			Id:        uuid.New(),
			ImageUrl:  imageUrl,
			ProductId: existingProduct.Id,
		})
	}

	newTags := make([]productModels.Tag, 0, len(productReq.Tags))

	for _, tagName := range productReq.Tags {
		foundTag := false
		for _, existingTag := range tags {
			if existingTag.TagName == tagName {
				newTags = append(newTags, productModels.Tag{
					Id:      existingTag.Id,
					TagName: tagName,
				})
				foundTag = true
				break
			}
		}

		if !foundTag {
			newTags = append(newTags, productModels.Tag{
				Id:      uuid.New(),
				TagName: tagName,
			})
		}
	}

	existingProduct.Name = productReq.Name
	existingProduct.Description = productReq.Description
	existingProduct.Price = productReq.Price
	existingProduct.Category = productReq.Category
	existingProduct.Images = images
	existingProduct.Tags = newTags
	existingProduct.UpdatedAt = time.Now()
	existingProduct.DeletedAt = nil
	return existingProduct

}

type ProductPatchRequest struct {
	Name        *string   `json:"name"`
	Description *string   `json:"description"`
	Price       *float64  `json:"price"`
	Category    *string   `json:"category"`
	Images      *[]string `json:"images"`
	Tags        *[]string `json:"tags"`
}

func PatchDtoToProduct(existingProduct *productModels.Product, productReq ProductPatchRequest, tags []productModels.Tag) *productModels.Product {

	if productReq.Name != nil {
		existingProduct.Name = *productReq.Name
	}

	if productReq.Description != nil {
		existingProduct.Description = *productReq.Description
	}

	if productReq.Price != nil {
		existingProduct.Price = *productReq.Price
	}

	if productReq.Category != nil {
		existingProduct.Category = *productReq.Category
	}

	if productReq.Images != nil {

		images := make([]productModels.ProductsImage, 0, len(*productReq.Images))

		for _, imageUrl := range *productReq.Images {
			images = append(images, productModels.ProductsImage{
				Id:        uuid.New(),
				ImageUrl:  imageUrl,
				ProductId: existingProduct.Id,
			})
		}

		existingProduct.Images = images

	}

	if productReq.Tags != nil {

		newTags := make([]productModels.Tag, 0, len(*productReq.Tags))

		for _, tagName := range *productReq.Tags {
			foundTag := false
			for _, existingTag := range tags {
				if existingTag.TagName == tagName {
					newTags = append(newTags, productModels.Tag{
						Id:      existingTag.Id,
						TagName: tagName,
					})
					foundTag = true
					break
				}
			}

			if !foundTag {
				newTags = append(newTags, productModels.Tag{
					Id:      uuid.New(),
					TagName: tagName,
				})
			}
		}

		existingProduct.Tags = newTags

	}

	existingProduct.UpdatedAt = time.Now()

	existingProduct.DeletedAt = nil

	return existingProduct

}

type IdsToDelete struct {
	Ids []uuid.UUID `json:"ids"`
}

type ProductQueryParams struct {
	Limit     int     `form:"limit"`
	Offset    int     `form:"offset"`
	Category  string  `form:"category"`
	MinPrice  float64 `form:"minPrice"`
	MaxPrice  float64 `form:"maxPrice"`
	SortBy    string  `form:"sortBy"`
	SortOrder string  `form:"sortOrder"`
	Search    string  `form:"search"`
}

type ProductsImageDto struct {
	Id        uuid.UUID `json:"id"`
	ImageUrl  string    `json:"image_url"`
	ProductId uuid.UUID `json:"product_id"`
}

type TagDto struct {
	Id      uuid.UUID `json:"id"`
	TagName string    `json:"tag_name"`
}

func ProductImagesToDto(product *productModels.Product) []ProductsImageDto {

	var productImages []ProductsImageDto

	for _, image := range product.Images {
		productImages = append(productImages, ProductsImageDto{Id: image.Id, ImageUrl: image.ImageUrl, ProductId: image.ProductId})
	}

	return productImages
}

func ProductTagsToDto(product *productModels.Product) []TagDto {

	var productTags []TagDto

	for _, tag := range product.Tags {
		productTags = append(productTags, TagDto{Id: tag.Id, TagName: tag.TagName})
	}

	return productTags
}

func ProductImagesToDtoString(product *productModels.Product) []string {

	var productImages []string

	for _, image := range product.Images {
		productImages = append(productImages, image.ImageUrl)
	}

	return productImages
}

func ProductTagsToDtoString(product *productModels.Product) []string {

	var productTags []string

	for _, tag := range product.Tags {
		productTags = append(productTags, tag.TagName)
	}

	return productTags
}

type ProductCatalogEventDto struct {
	Id          uuid.UUID `json:"id"`
	VendorId    uuid.UUID `json:"vendor_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       string    `json:"price"`
	Category    string    `json:"category"`
	Images      []string  `json:"images"`
	Tags        []string  `json:"tags"`
	Quantity    int       `json:"quantity"`
}

func ProductToEventDto(product *productModels.Product) ProductCatalogEventDto {
	return ProductCatalogEventDto{
		Id:          product.Id,
		VendorId:    product.VendorId,
		Name:        product.Name,
		Description: product.Description,
		Price:       strconv.FormatFloat(product.Price, 'f', -1, 64),
		Category:    product.Category,
		Images:      ProductImagesToDtoString(product),
		Tags:        ProductTagsToDtoString(product),
		Quantity:    product.Quantity,
	}
}

type ProductUpdatedCatalogEvent struct {
	EventId uuid.UUID              `json:"id"`
	Product ProductCatalogEventDto `json:"product"`
}

type QuantityStatusEvent struct {
	EventId   uuid.UUID `json:"id"`
	Type      string    `json:"type"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func ProductToOutbox(product *productModels.Product, eventType, exchange string) (*models.Outbox, error) {

	event := ProductUpdatedCatalogEvent{
		EventId: uuid.New(),
		Product: ProductToEventDto(product),
	}

	payload, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	return &models.Outbox{
		Id:          uuid.New(),
		Exchange:    exchange,
		EventType:   eventType,
		Payload:     payload,
		CreatedAt:   time.Now(),
		Processed:   false,
		ProcessedAt: time.Time{},
	}, nil
}

type DeletedProductEventDto struct {
	EventId uuid.UUID   `json:"event_id"`
	Ids     []uuid.UUID `json:"deleted_ids"`
}

func DeletedProductToOutbox(deletedIds []uuid.UUID, eventType, exchange string) (*models.Outbox, error) {

	event := DeletedProductEventDto{
		EventId: uuid.New(),
		Ids:     deletedIds,
	}

	payload, err := json.Marshal(event)

	if err != nil {
		return nil, err
	}

	return &models.Outbox{
		Id:          uuid.New(),
		Exchange:    exchange,
		EventType:   eventType,
		Payload:     payload,
		CreatedAt:   time.Now(),
		Processed:   false,
		ProcessedAt: time.Time{},
	}, nil
}

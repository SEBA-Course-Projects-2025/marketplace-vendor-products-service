package dtos

import (
	"dev-vendor/product-service/internal/products/domain/models"
	"github.com/google/uuid"
	"time"
)

type OneProductResponse struct {
	Id          uuid.UUID              `json:"id"`
	VendorId    uuid.UUID              `json:"vendor_id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description"`
	Price       float64                `json:"price"`
	Category    string                 `json:"category"`
	Images      []models.ProductsImage `json:"images"`
	Tags        []models.Tag           `json:"tags"`
	Quantity    int                    `json:"quantity"`
}

func ProductToDto(product models.Product) OneProductResponse {
	return OneProductResponse{
		Id:          product.Id,
		VendorId:    product.VendorId,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Category:    product.Category,
		Images:      product.Images,
		Tags:        product.Tags,
		Quantity:    product.Quantity,
	}
}

type GetProductsResponse struct {
	Id       uuid.UUID            `json:"id"`
	VendorId uuid.UUID            `json:"vendor_id"`
	Name     string               `json:"name"`
	Price    float64              `json:"price"`
	Category string               `json:"category"`
	Image    models.ProductsImage `json:"image"`
	Quantity int                  `json:"quantity"`
}

func ProductsToDto(products []models.Product) []GetProductsResponse {

	var productsResponse []GetProductsResponse

	for _, product := range products {

		var image models.ProductsImage

		if len(product.Images) > 0 {
			image = product.Images[0]
		} else {
			image = models.ProductsImage{}
		}

		productResponse := GetProductsResponse{
			Id:       product.Id,
			VendorId: product.VendorId,
			Name:     product.Name,
			Price:    product.Price,
			Category: product.Category,
			Image:    image,
			Quantity: product.Quantity,
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

func PostDtoToProduct(productReq ProductRequest, vendorId uuid.UUID) models.Product {

	productId := uuid.New()

	images := make([]models.ProductsImage, 0, len(productReq.Images))

	for _, imageUrl := range productReq.Images {
		images = append(images, models.ProductsImage{
			Id:        uuid.New(),
			ImageUrl:  imageUrl,
			ProductId: productId,
		})
	}

	tags := make([]models.Tag, 0, len(productReq.Tags))

	for _, tag := range productReq.Tags {
		tags = append(tags, models.Tag{
			Id:      uuid.New(),
			TagName: tag,
		})
	}

	product := models.Product{
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

	return product
}

func UpdateProductWithDto(existingProduct *models.Product, productReq ProductRequest) *models.Product {

	images := make([]models.ProductsImage, 0, len(productReq.Images))

	for _, imageUrl := range productReq.Images {
		for _, image := range existingProduct.Images {
			if image.ImageUrl == imageUrl {
				images = append(images, models.ProductsImage{
					Id:        image.Id,
					ImageUrl:  imageUrl,
					ProductId: existingProduct.Id,
				})
				break
			}
		}
	}

	tags := make([]models.Tag, 0, len(productReq.Tags))

	for _, tagName := range productReq.Tags {
		for _, tag := range existingProduct.Tags {
			if tag.TagName == tagName {
				tags = append(tags, models.Tag{
					Id:      tag.Id,
					TagName: tagName,
				})
				break
			}
		}
	}

	existingProduct.Name = productReq.Name
	existingProduct.Description = productReq.Description
	existingProduct.Price = productReq.Price
	existingProduct.Category = productReq.Category
	existingProduct.Images = images
	existingProduct.Tags = tags
	existingProduct.UpdatedAt = time.Now()
	return existingProduct

}

type ProductPatchRequest struct {
	Name        *string                 `json:"name"`
	Description *string                 `json:"description"`
	Price       *float64                `json:"price"`
	Category    *string                 `json:"category"`
	Images      *[]models.ProductsImage `json:"images"`
	Tags        *[]models.Tag           `json:"tags"`
}

func PatchDtoToProduct(existingProduct *models.Product, productReq ProductPatchRequest) *models.Product {

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

		updatedImages := make([]models.ProductsImage, 0, len(*productReq.Images))
		for _, patchImg := range *productReq.Images {
			for _, existImg := range existingProduct.Images {
				if patchImg.ImageUrl == existImg.ImageUrl {
					updatedImages = append(updatedImages, existImg)
					break
				}
			}
		}

		existingProduct.Images = *productReq.Images
	}

	if productReq.Tags != nil {

		updatedTags := make([]models.Tag, 0, len(*productReq.Tags))
		for _, patchTag := range *productReq.Tags {
			for _, existTag := range existingProduct.Tags {
				if patchTag.TagName == existTag.TagName {
					updatedTags = append(updatedTags, existTag)
					break
				}
			}
		}

		existingProduct.Tags = *productReq.Tags
	}

	existingProduct.UpdatedAt = time.Now()

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

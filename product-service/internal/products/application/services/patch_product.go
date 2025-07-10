package services

import (
	"context"
	eventDomain "dev-vendor/product-service/internal/event/domain"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func PatchProduct(ctx context.Context, repo domain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, id uuid.UUID, productReq dtos.ProductPatchRequest, vendorId uuid.UUID) (dtos.OneProductResponse, error) {

	ctx, span := tracer.Tracer.Start(ctx, "PatchProductById")
	defer span.End()

	var productResponse dtos.OneProductResponse

	if err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txProductRepo := repo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		existingProduct, err := txProductRepo.FindById(ctx, id)

		if err != nil {
			return err
		}

		if productReq.Images != nil {
			if err := txProductRepo.DeleteProductImages(ctx, existingProduct); err != nil {
				return err
			}
		}

		if productReq.Tags != nil {
			if err := txProductRepo.DeleteProductTags(ctx, existingProduct); err != nil {
				return err
			}
		}

		tags, err := txProductRepo.FindAllTags(ctx)

		if err != nil {
			return err
		}

		existingProduct = dtos.PatchDtoToProduct(existingProduct, productReq, tags)

		existingProduct.VendorId = vendorId

		existingProduct, err = txProductRepo.Patch(ctx, existingProduct)

		if err != nil {
			return err
		}

		outbox, err := dtos.ProductToOutbox(existingProduct, "product.updated.catalog", "product.catalog.events")

		if err != nil {
			return err
		}

		err = txEventRepo.CreateOutboxRecord(ctx, outbox)

		if err != nil {
			return err
		}

		productResponse = dtos.ProductToDto(existingProduct)
		return nil

	}); err != nil {
		return dtos.OneProductResponse{}, err
	}

	return productResponse, nil
}

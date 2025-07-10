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

func PutProduct(ctx context.Context, repo domain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, id uuid.UUID, productReq dtos.ProductRequest, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "PutProductById")
	defer span.End()

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txProductRepo := repo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		existingProduct, err := txProductRepo.FindById(ctx, id)

		if err != nil {
			return err
		}

		if err := txProductRepo.DeleteProductImages(ctx, existingProduct); err != nil {
			return err
		}

		if err := txProductRepo.DeleteProductTags(ctx, existingProduct); err != nil {
			return err
		}

		tags, err := txProductRepo.FindAllTags(ctx)

		if err != nil {
			return err
		}

		existingProduct = dtos.UpdateProductWithDto(existingProduct, productReq, tags)
		existingProduct.VendorId = vendorId

		err = txProductRepo.Update(ctx, existingProduct)

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

		return nil
	})

}

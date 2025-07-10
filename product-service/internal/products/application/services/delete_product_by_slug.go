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

func DeleteProductBySlug(ctx context.Context, repo domain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, slug string, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteOneProductBySlug")
	defer span.End()

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txProductRepo := repo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		product, err := txProductRepo.FindBySlug(ctx, slug, vendorId)

		if err != nil {
			return err
		}

		if err := txProductRepo.DeleteBySlug(ctx, slug, vendorId); err != nil {
			return err
		}

		outbox, err := dtos.DeletedProductToOutbox([]uuid.UUID{product.Id}, "product.deleted.catalog", "product.catalog.events")

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

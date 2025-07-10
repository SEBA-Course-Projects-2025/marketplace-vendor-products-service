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

func DeleteProductById(ctx context.Context, repo domain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, id uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteOneProductById")
	defer span.End()

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txProductRepo := repo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		if err := txProductRepo.DeleteById(ctx, id, vendorId); err != nil {
			return err
		}

		outbox, err := dtos.DeletedProductToOutbox([]uuid.UUID{id}, "product.deleted.catalog", "product.catalog.events")

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

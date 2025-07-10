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

func DeleteManyProducts(ctx context.Context, repo domain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, ids []uuid.UUID, vendorId uuid.UUID) error {

	ctx, span := tracer.Tracer.Start(ctx, "DeleteAllProducts")
	defer span.End()

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		txProductRepo := repo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		if err := txProductRepo.DeleteMany(ctx, ids, vendorId); err != nil {
			return err
		}

		outbox, err := dtos.DeletedProductToOutbox(ids, "product.deleted.catalog", "product.catalog.events")

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

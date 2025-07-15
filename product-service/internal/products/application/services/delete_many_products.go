package services

import (
	"context"
	eventDomain "dev-vendor/product-service/internal/event/domain"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/shared/utils"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func DeleteManyProducts(ctx context.Context, repo domain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, ids []uuid.UUID, vendorId uuid.UUID) error {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
		"ids":      ids,
	}).Info("Starting DeleteManyProducts application service")

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
			return utils.ErrorHandler(err, err.Error())
		}

		err = txEventRepo.CreateOutboxRecord(ctx, outbox)

		if err != nil {
			return utils.ErrorHandler(err, err.Error())
		}

		logrus.WithFields(logrus.Fields{
			"vendorId": vendorId,
			"ids":      ids,
		}).Info("Successfully soft deleted products")
		return nil
	})

}

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

func PutProductBySlug(ctx context.Context, repo domain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, slug string, productReq dtos.ProductRequest, vendorId uuid.UUID) error {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
		"slug":     slug,
	}).Info("Starting PutProductBySlug application service")

	ctx, span := tracer.Tracer.Start(ctx, "PutProductBySlug")
	defer span.End()

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txProductRepo := repo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		existingProduct, err := txProductRepo.FindBySlug(ctx, slug, vendorId)

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
			return utils.ErrorHandler(err, err.Error())
		}

		err = txEventRepo.CreateOutboxRecord(ctx, outbox)

		if err != nil {
			return err
		}

		logrus.WithFields(logrus.Fields{
			"vendorId": vendorId,
			"slug":     slug,
		}).Info("Successfully fully modified product by slug")

		return nil
	})

}

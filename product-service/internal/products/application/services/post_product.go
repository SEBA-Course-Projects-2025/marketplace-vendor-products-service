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

func PostProduct(ctx context.Context, repo domain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, productReq dtos.ProductRequest, vendorId uuid.UUID) (dtos.OneProductResponse, error) {

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Starting PostProduct application service")

	ctx, span := tracer.Tracer.Start(ctx, "PostProduct")
	defer span.End()

	var productResponse dtos.OneProductResponse

	if err := db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txProductRepo := repo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		newProduct, err := dtos.PostDtoToProduct(productReq, vendorId)
		if err != nil {
			return err
		}

		product, err := txProductRepo.Create(ctx, &newProduct, vendorId)

		if err != nil {
			return err
		}

		outbox, err := dtos.ProductToOutbox(product, "product.created.catalog", "product.catalog.events")

		if err != nil {
			return utils.ErrorHandler(err, err.Error())
		}

		err = txEventRepo.CreateOutboxRecord(ctx, outbox)

		if err != nil {
			return err
		}

		productResponse = dtos.ProductToDto(product)

		return nil
	}); err != nil {
		return dtos.OneProductResponse{}, err
	}

	logrus.WithFields(logrus.Fields{
		"vendorId": vendorId,
	}).Info("Successfully created product")

	return productResponse, nil

}

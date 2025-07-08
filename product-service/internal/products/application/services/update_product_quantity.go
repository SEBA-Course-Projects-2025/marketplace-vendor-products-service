package services

import (
	"context"
	eventDomain "dev-vendor/product-service/internal/event/domain"
	"dev-vendor/product-service/internal/event/domain/models"
	"dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/dtos"
	"dev-vendor/product-service/internal/shared/tracer"
	"github.com/google/uuid"
)

func UpdateProductQuantity(ctx context.Context, repo domain.ProductRepository, eventRepo eventDomain.EventRepository, productId uuid.UUID, quantity int, exchange string) error {

	ctx, span := tracer.Tracer.Start(ctx, "UpdateProductQuantity")
	defer span.End()

	product, err := repo.FindById(ctx, productId)

	if err != nil {
		return err
	}

	product.Quantity = quantity

	if err := repo.Update(ctx, product); err != nil {
		return err
	}

	var outbox *models.Outbox

	outbox, err = dtos.ProductToOutbox(product, "product.updated.catalog", "product.catalog.events")
	if err != nil {
		return err
	}

	err = eventRepo.CreateOutboxRecord(ctx, outbox)

	if err != nil {
		return err
	}

	return nil

}

package domain

import (
	"context"
	"dev-vendor/product-service/internal/event/domain/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EventRepository interface {
	CreateOutboxRecord(ctx context.Context, outbox *models.Outbox) error
	FetchUnprocessed(ctx context.Context) ([]models.Outbox, error)
	MarkProcessed(ctx context.Context, id uuid.UUID) error
	CheckProcessedMessage(ctx context.Context, id uuid.UUID) (bool, error)
	CreateProcessedMessage(ctx context.Context, id uuid.UUID) error
	Transaction(fn func(txRepo EventRepository) error) error
	WithTx(tx *gorm.DB) EventRepository
}

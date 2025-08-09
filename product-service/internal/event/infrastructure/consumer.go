package infrastructure

import (
	"context"
	"dev-vendor/product-service/internal/stocks/application/services"
	"dev-vendor/product-service/internal/stocks/dtos"
	"dev-vendor/product-service/internal/stocks/interfaces/handlers"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"log"
)

type Consumer struct {
	AMQPChannel  *amqp.Channel
	StockHandler *handlers.StockHandler
}

func NewConsumer(channel *amqp.Channel, stockHandler *handlers.StockHandler) *Consumer {
	return &Consumer{
		AMQPChannel:  channel,
		StockHandler: stockHandler,
	}
}

func (c *Consumer) StartConsuming(ctx context.Context, queueName string) error {

	msgs, err := c.AMQPChannel.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-msgs:
				if !ok {
					return
				}

				msgId, err := uuid.Parse(msg.MessageId)
				if err != nil {
					log.Printf("Invalid message id: %v", err)
					_ = msg.Nack(false, false)
					continue
				}

				processed, err := c.StockHandler.EventRepo.CheckProcessedMessage(ctx, msgId)
				if err != nil {
					log.Printf("Error checking processed message: %v", err)
					_ = msg.Nack(false, false)
					continue
				}
				if processed {
					_ = msg.Ack(false)
					continue
				}

				if err = c.processMessage(msg, ctx); err != nil {
					log.Printf("Error processing new message: %v", err)
					_ = msg.Nack(false, false)
					continue
				}

				if err = c.StockHandler.EventRepo.CreateProcessedMessage(ctx, msgId); err != nil {
					log.Printf("Error adding processed message: %v", err)
					_ = msg.Nack(false, false)
					continue
				}

				_ = msg.Ack(false)
			}
		}
	}()

	return nil

}

func (c *Consumer) processMessage(msg amqp.Delivery, ctx context.Context) error {

	eventType := msg.Type

	if eventType == "" {
		return fmt.Errorf("event type not found")
	}

	switch eventType {

	case "vendor.check.product.quantity":

		var quantityDto dtos.OrderCreatedEventDto
		if err := json.Unmarshal(msg.Body, &quantityDto); err != nil {
			return err
		}

		if err := c.StockHandler.ReduceStockProductQuantityHandler(ctx, quantityDto); err != nil {
			return err
		}

	case "vendor.cancel.product.order":

		var canceledOrderProducts []dtos.CanceledOrderItemDto

		if err := json.Unmarshal(msg.Body, &canceledOrderProducts); err != nil {
			return err
		}

		if err := services.ReturnCanceledQuantity(ctx, c.StockHandler.StockRepo, c.StockHandler.ProductRepo, c.StockHandler.EventRepo, c.StockHandler.Db, canceledOrderProducts); err != nil {
			return err
		}

	default:
		log.Printf("Unknown event type: %s", eventType)
		return nil
	}

	return nil
}

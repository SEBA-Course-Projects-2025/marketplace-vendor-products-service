package services

import (
	"context"
	eventDomain "dev-vendor/product-service/internal/event/domain"
	"dev-vendor/product-service/internal/products/application/services"
	productDomain "dev-vendor/product-service/internal/products/domain"
	"dev-vendor/product-service/internal/products/domain/productModels"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/domain"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"dev-vendor/product-service/internal/stocks/dtos"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func ReduceStockProductQuantity(ctx context.Context, stockRepo domain.StockRepository, productRepo productDomain.ProductRepository, eventRepo eventDomain.EventRepository, db *gorm.DB, quantityReq dtos.OrderCreatedEventDto) error {

	ctx, span := tracer.Tracer.Start(ctx, "ReduceStockProductQuantity")
	defer span.End()

	var products []productModels.Product
	var items []dtos.OrderItemDto
	stockIds := make(map[uuid.UUID]uuid.UUID)

	return db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		txStockRepo := stockRepo.WithTx(tx)
		txProductRepo := productRepo.WithTx(tx)
		txEventRepo := eventRepo.WithTx(tx)

		for _, product := range quantityReq.Items {

			checkProduct, err := txProductRepo.FindById(ctx, product.ProductId)

			if err != nil {
				return err
			}

			existingProductStocksQuantities, err := txStockRepo.FindProductStocksQuantities(ctx, checkProduct.Id)

			if err != nil {
				return err
			}

			remainToReduce := product.Quantity

			for _, q := range existingProductStocksQuantities {

				stockIds[checkProduct.Id] = q.StockId

				if remainToReduce == 0 {
					break
				}

				toReduce := q.Quantity
				if toReduce > remainToReduce {
					toReduce = remainToReduce
				}

				existingStock, err := txStockRepo.FindById(ctx, q.StockId)
				if err != nil {
					return err
				}

				var existingStockProduct *models.StocksProduct
				for i := range existingStock.StocksProducts {
					if existingStock.StocksProducts[i].ProductId == checkProduct.Id {
						existingStockProduct = &existingStock.StocksProducts[i]
						break
					}
				}

				if existingStockProduct == nil {
					return fmt.Errorf("stock product is not found")
				}

				existingStockProduct.Quantity -= toReduce

				if existingStockProduct.Quantity == 0 {
					if err := txStockRepo.DeleteStockProductById(ctx, existingStockProduct.StockId, existingStockProduct.ProductId, checkProduct.VendorId); err != nil {
						return err
					}
				} else {
					if _, err = txStockRepo.PatchStockProductId(ctx, existingStockProduct); err != nil {
						return err
					}
				}

				remainToReduce -= toReduce

			}

			if remainToReduce > 0 {

				resp := dtos.OrderCreatedEventResponseDto{
					EventId:    quantityReq.EventId,
					OrderId:    quantityReq.OrderId,
					CustomerId: quantityReq.CustomerId,
					Items:      nil,
					TotalPrice: quantityReq.TotalPrice,
					Status:     "declined",
				}

				outbox, err := dtos.QuantityStatusToOutbox(resp, "vendor.product.quantity.checked", "vendor.product.events")
				if err != nil {
					return err
				}
				err = eventRepo.CreateOutboxRecord(ctx, outbox)

				if err != nil {
					return err
				}

				return fmt.Errorf("not enough stock product quantities to reduce")
			}

			quantitySum, err := GetQuantitySum(ctx, txStockRepo, checkProduct.Id)

			if err != nil {
				return err
			}

			if err := services.UpdateProductQuantity(ctx, txProductRepo, txEventRepo, checkProduct.Id, quantitySum, "vendor.product.events"); err != nil {
				return err
			}

			products = append(products, *checkProduct)
			items = append(items, product)
		}

		itemsToSend := dtos.OrderItemsToEventResponseDto(products, items, stockIds)

		resp := dtos.OrderCreatedEventResponseDto{
			EventId:    quantityReq.EventId,
			OrderId:    quantityReq.OrderId,
			CustomerId: quantityReq.CustomerId,
			Items:      itemsToSend,
			TotalPrice: quantityReq.TotalPrice,
			Status:     "confirmed",
		}

		outbox, err := dtos.QuantityStatusToOutbox(resp, "vendor.product.quantity.checked", "vendor.product.events")
		if err != nil {
			return err
		}
		err = eventRepo.CreateOutboxRecord(ctx, outbox)

		if err != nil {
			return err
		}

		return nil
	})
}

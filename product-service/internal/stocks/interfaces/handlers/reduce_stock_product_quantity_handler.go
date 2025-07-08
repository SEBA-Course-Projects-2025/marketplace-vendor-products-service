package handlers

import (
	"context"
	"dev-vendor/product-service/internal/shared/tracer"
	"dev-vendor/product-service/internal/stocks/application/services"
	"dev-vendor/product-service/internal/stocks/dtos"
)

func (h *StockHandler) ReduceStockProductQuantityHandler(ctx context.Context, eventDto dtos.OrderCreatedEventDto) error {

	ctx, span := tracer.Tracer.Start(ctx, "ReduceStockProductQuantityHandler")
	defer span.End()

	return services.ReduceStockProductQuantity(ctx, h.StockRepo, h.ProductRepo, h.EventRepo, h.Db, eventDto)

}

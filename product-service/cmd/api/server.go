package main

import (
	"context"
	eventInfrastructure "dev-vendor/product-service/internal/event/infrastructure"
	eventRepository "dev-vendor/product-service/internal/event/infrastructure/repository"
	productRepository "dev-vendor/product-service/internal/products/infrastructure/repository"
	productHandlers "dev-vendor/product-service/internal/products/interfaces/handlers"
	"dev-vendor/product-service/internal/shared/amqp"
	"dev-vendor/product-service/internal/shared/db"
	mainHandler "dev-vendor/product-service/internal/shared/handler"
	"dev-vendor/product-service/internal/shared/router"
	"dev-vendor/product-service/internal/shared/tracer"
	stockRepository "dev-vendor/product-service/internal/stocks/infrastructure/repository"
	stockHandlers "dev-vendor/product-service/internal/stocks/interfaces/handlers"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

// @title Product Service API
// @version 1.0
// @description API for managing products and stocks for vendors.

// @schemes https
// @host marketplace-vendor-products-service.onrender.com
// @BasePath /api
func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUsed, err := db.ConnectDb()
	if err != nil {
		log.Fatalln(err)
	}

	newTracer := tracer.InitTracer()
	defer func() {
		if err := newTracer(context.Background()); err != nil {
			log.Fatalf("Error shutting down tracer: %v", err)
		}
	}()

	productRepo := productRepository.New(dbUsed)
	stockRepo := stockRepository.New(dbUsed)
	eventRepo := eventRepository.New(dbUsed)

	amqpConfig, err := amqp.ConnectAMQP()
	if err != nil {
		log.Fatalln(err)
	}

	outboxPoller := eventInfrastructure.NewOutboxPoller(eventRepo, amqpConfig.Channel, time.Second*2)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		if err := outboxPoller.StartPolling(ctx, amqpConfig.ConfirmationChannel); err != nil {
			log.Fatalln(err)
		}
	}()

	sharedHandler := &mainHandler.Handler{
		ProductRepo: productRepo,
		StockRepo:   stockRepo,
		EventRepo:   eventRepo,
		Db:          dbUsed,
	}

	productHandler := &productHandlers.ProductHandler{
		Handler: sharedHandler,
	}

	stockHandler := &stockHandlers.StockHandler{
		Handler: sharedHandler,
	}

	consumer := eventInfrastructure.NewConsumer(amqpConfig.Channel, stockHandler)

	queues := []string{
		"vendor.check.product.quantity",
		"vendor.cancel.product.order",
	}

	for _, queue := range queues {
		go func(q string) {
			if err := consumer.StartConsuming(ctx, queue); err != nil {
				log.Fatalf("Consumer error: %v", err)
			}
		}(queue)
	}

	mainRouter := router.SetUpRouter(productHandler, stockHandler)

	port := os.Getenv("API_PORT")

	if port == "" {
		port = "8080"
	}

	fmt.Println(time.Now())

	if err := mainRouter.Run(":" + port); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}

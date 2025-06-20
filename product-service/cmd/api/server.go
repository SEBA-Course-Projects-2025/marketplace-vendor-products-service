package main

import (
	productRepository "dev-vendor/product-service/internal/products/infrastructure/repository"
	productHandlers "dev-vendor/product-service/internal/products/interfaces/handlers"
	"dev-vendor/product-service/internal/shared/db"
	mainHandler "dev-vendor/product-service/internal/shared/handler"
	"dev-vendor/product-service/internal/shared/router"
	stockRepository "dev-vendor/product-service/internal/stocks/infrastructure/repository"
	stockHandlers "dev-vendor/product-service/internal/stocks/interfaces/handlers"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	dbUsed, err := db.ConnectDb()
	if err != nil {
		log.Fatalln(err)
	}

	productRepo := productRepository.New(dbUsed)
	stockRepo := stockRepository.New(dbUsed)

	sharedHandler := &mainHandler.Handler{
		ProductRepo: productRepo,
		StockRepo:   stockRepo,
		Db:          dbUsed,
	}

	productHandler := &productHandlers.ProductHandler{
		Handler: sharedHandler,
	}

	stockHandler := &stockHandlers.StockHandler{
		Handler: sharedHandler,
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

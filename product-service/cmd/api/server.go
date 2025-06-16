package main

import (
	productRepository "dev-vendor/product-service/internal/products/infrastucture/repository"
	productHandlers "dev-vendor/product-service/internal/products/interfaces/handlers"
	"dev-vendor/product-service/internal/shared/db"
	mainHandler "dev-vendor/product-service/internal/shared/handler"
	"dev-vendor/product-service/internal/shared/router"
	stockRepository "dev-vendor/product-service/internal/stocks/infrastucture/repository"
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

	productRepo := &productRepository.GormProductRepository{Db: dbUsed}
	stockRepo := &stockRepository.GormStockRepository{Db: dbUsed}

	productHandler := &productHandlers.ProductHandler{
		Handler: &mainHandler.Handler{
			ProductRepo: productRepo,
			StockRepo:   stockRepo,
		},
	}

	stockHandler := &stockHandlers.StockHandler{
		Handler: &mainHandler.Handler{
			StockRepo:   stockRepo,
			ProductRepo: productRepo,
		},
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

package main

import (
	"customer-mock-api-gateway/customer-mock-api-gateway/internal/api/routers"
	"log"
	"os"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "8081"
	}

	r := routers.SetUpRouter()
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

}

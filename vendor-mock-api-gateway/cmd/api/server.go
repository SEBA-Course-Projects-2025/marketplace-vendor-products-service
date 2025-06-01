package main

import (
	"log"
	"os"
	"vendor-mock-api-gateway/vendor-mock-api-gateway/internal/api/routers"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routers.SetUpRouter()
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}

}

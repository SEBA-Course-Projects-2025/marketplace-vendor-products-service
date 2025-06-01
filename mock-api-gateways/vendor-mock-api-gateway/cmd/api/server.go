package main

import (
	"log"
	"mock-api-gateway/mock-api-gateways/vendor-mock-api-gateway/internal/api/routers"
)

func main() {
	r := routers.SetUpRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

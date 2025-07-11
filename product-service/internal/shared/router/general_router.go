package router

import (
	_ "dev-vendor/docs"
	productInterfaces "dev-vendor/product-service/internal/products/interfaces"
	productsHandlers "dev-vendor/product-service/internal/products/interfaces/handlers"
	"dev-vendor/product-service/internal/shared/middlewares"
	stocksInterfaces "dev-vendor/product-service/internal/stocks/interfaces"
	stockHandlers "dev-vendor/product-service/internal/stocks/interfaces/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"log"
	"os"
)

func SetUpRouter(productHandler *productsHandlers.ProductHandler, stockHandler *stockHandlers.StockHandler) *gin.Engine {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	r := gin.New()

	p := ginprometheus.NewPrometheus("vendor_product_service")
	p.MetricsPath = ""
	p.Use(r)

	r.GET("/metrics", gin.BasicAuth(gin.Accounts{
		os.Getenv("METRICS_ACCESS_USERNAME"): os.Getenv("METRICS_ACCESS_PASSWORD"),
	}), gin.WrapH(promhttp.Handler()))

	r.Use(otelgin.Middleware("vendor_product_service"))

	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-length"},
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api", middlewares.AuthMiddleware())

	productInterfaces.SetUpProductsRouter(api, productHandler)
	stocksInterfaces.SetUpStocksRouter(api, stockHandler)

	return r
}

package dtos

import (
	products_models "dev-vendor/product-service/internal/products/domain/models"
	"dev-vendor/product-service/internal/stocks/domain/models"
	"github.com/google/uuid"
	"time"
)

type OneStockResponse struct {
	Id           uuid.UUID             `json:"id"`
	VendorId     uuid.UUID             `json:"vendor_id"`
	DateSupplied time.Time             `json:"date_supplied"`
	Location     models.StocksLocation `json:"location"`
	Products     []StockProductInfo    `json:"products"`
}

type StockProductInfo struct {
	Id       uuid.UUID                     `json:"product_id"`
	Name     string                        `json:"name"`
	Quantity int                           `json:"quantity"`
	UnitCost float64                       `json:"unit_cost"`
	Image    products_models.ProductsImage `json:"image"`
}

func StockToDto(stock *models.Stock) OneStockResponse {

	products := make([]StockProductInfo, len(stock.StocksProducts))

	for i, stockProduct := range stock.StocksProducts {

		var image products_models.ProductsImage

		if len(stockProduct.Product.Images) > 0 {
			image = stockProduct.Product.Images[0]
		} else {
			image = products_models.ProductsImage{}
		}

		products[i] = StockProductInfo{
			Id:       stockProduct.Product.Id,
			Name:     stockProduct.Product.Name,
			Quantity: stockProduct.Quantity,
			UnitCost: stockProduct.UnitCost,
			Image:    image,
		}
	}

	return OneStockResponse{
		Id:           stock.Id,
		VendorId:     stock.VendorId,
		DateSupplied: stock.DateSupplied,
		Location:     stock.Location,
		Products:     products,
	}
}

type PostStockProductRequest struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  int       `json:"quantity"`
	UnitCost  float64   `json:"unit_cost"`
}

type StockRequest struct {
	DateSupplied time.Time                 `json:"date_supplied"`
	LocationId   uuid.UUID                 `json:"location_id"`
	Products     []PostStockProductRequest `json:"products"`
}

func PostStockRequestToStock(stockReq StockRequest, vendorId uuid.UUID) (*models.Stock, error) {

	stockId := uuid.New()
	var stocksProducts []models.StocksProduct

	for _, product := range stockReq.Products {
		stocksProducts = append(stocksProducts, models.StocksProduct{
			StockId:   stockId,
			ProductId: product.ProductId,
			Quantity:  product.Quantity,
			UnitCost:  product.UnitCost,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
	}

	stock := &models.Stock{
		Id:             uuid.New(),
		VendorId:       vendorId,
		DateSupplied:   stockReq.DateSupplied,
		LocationId:     stockReq.LocationId,
		StocksProducts: stocksProducts,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	return stock, nil
}

type PostStockResponse struct {
	Id           uuid.UUID `json:"id"`
	VendorId     uuid.UUID `json:"vendor_id"`
	DateSupplied time.Time `json:"date_supplied"`
	LocationId   uuid.UUID `json:"location_id"`
}

func PostStockToStockResponse(stock *models.Stock) PostStockResponse {
	return PostStockResponse{
		Id:           stock.Id,
		VendorId:     stock.VendorId,
		DateSupplied: stock.DateSupplied,
		LocationId:   stock.LocationId,
	}
}

type PutStockRequest struct {
	DateSupplied time.Time `json:"date_supplied"`
	LocationId   uuid.UUID `json:"location_id"`
}

func UpdateStockWithDto(existingStock *models.Stock, stockReq PutStockRequest, location *models.StocksLocation) *models.Stock {

	return &models.Stock{
		Id:             existingStock.Id,
		VendorId:       existingStock.VendorId,
		DateSupplied:   stockReq.DateSupplied,
		LocationId:     stockReq.LocationId,
		Location:       *location,
		StocksProducts: existingStock.StocksProducts,
		CreatedAt:      existingStock.CreatedAt,
		UpdatedAt:      time.Now(),
	}

}

type PutStockProductRequest struct {
	Quantity int     `json:"quantity"`
	UnitCost float64 `json:"unit_cost"`
}

func UpdateStockProductWithDto(existingStockProduct *models.StocksProduct, stockProductReq PutStockProductRequest) *models.StocksProduct {

	return &models.StocksProduct{
		StockId:   existingStockProduct.StockId,
		ProductId: existingStockProduct.ProductId,
		Quantity:  stockProductReq.Quantity,
		UnitCost:  stockProductReq.UnitCost,
		CreatedAt: existingStockProduct.CreatedAt,
		UpdatedAt: time.Now(),
		Stock:     existingStockProduct.Stock,
		Product:   existingStockProduct.Product,
	}

}

type StockPatchRequest struct {
	DateSupplied *time.Time `json:"date_supplied,omitempty"`
	LocationId   *uuid.UUID `json:"location_id,omitempty"`
}

func ModifyStockWithDto(existingStock *models.Stock, stockReq StockPatchRequest, location *models.StocksLocation) *models.Stock {

	if stockReq.DateSupplied != nil {
		existingStock.DateSupplied = *stockReq.DateSupplied
	}

	if stockReq.LocationId != nil {
		existingStock.LocationId = *stockReq.LocationId
		existingStock.Location = *location
	}

	return existingStock

}

type PatchStockProductRequest struct {
	Quantity *int     `json:"quantity,omitempty"`
	UnitCost *float64 `json:"unit_cost,omitempty"`
}

func ModifyStockProductWithDto(existingStockProduct *models.StocksProduct, stockProductReq PatchStockProductRequest) *models.StocksProduct {

	if stockProductReq.Quantity != nil && *stockProductReq.Quantity >= 1 {
		existingStockProduct.Quantity = *stockProductReq.Quantity
	}

	if stockProductReq.UnitCost != nil && *stockProductReq.UnitCost > 0 {
		existingStockProduct.UnitCost = *stockProductReq.UnitCost
	}

	return existingStockProduct

}

func StocksProductToStockProductInfo(updatedStockProduct *models.StocksProduct) StockProductInfo {

	return StockProductInfo{
		Id:       updatedStockProduct.ProductId,
		Name:     updatedStockProduct.Product.Name,
		Quantity: updatedStockProduct.Product.Quantity,
		UnitCost: updatedStockProduct.UnitCost,
		Image:    updatedStockProduct.Product.Images[0],
	}

}

type PatchStockManyProductsRequest struct {
	ProductId uuid.UUID `json:"product_id"`
	Quantity  *int      `json:"quantity,omitempty"`
	UnitCost  *float64  `json:"unit_cost,omitempty"`
}

func ModifyStockManyProductsWithDto(existingStockProduct *models.StocksProduct, stockProductReq PatchStockManyProductsRequest) *models.StocksProduct {

	if stockProductReq.Quantity != nil && *stockProductReq.Quantity >= 1 {
		existingStockProduct.Quantity = *stockProductReq.Quantity
	}

	if stockProductReq.UnitCost != nil && *stockProductReq.UnitCost > 0 {
		existingStockProduct.UnitCost = *stockProductReq.UnitCost
	}

	return existingStockProduct

}

type StockQueryParams struct {
	Limit      int    `form:"limit"`
	Offset     int    `form:"offset"`
	LocationId string `form:"location_id"`
	SortBy     string `form:"sortBy"`
	SortOrder  string `form:"sortOrder"`
}

type GetStocksResponse struct {
	Id           uuid.UUID `json:"id"`
	DateSupplied time.Time `json:"date_supplied"`
	LocationId   uuid.UUID `json:"location_id"`
}

func StocksToDto(stocks []models.Stock) []GetStocksResponse {

	var getStocksResponses []GetStocksResponse

	for _, stock := range stocks {
		getStock := GetStocksResponse{
			Id:           stock.Id,
			DateSupplied: stock.DateSupplied,
			LocationId:   stock.LocationId,
		}
		getStocksResponses = append(getStocksResponses, getStock)
	}

	return getStocksResponses

}

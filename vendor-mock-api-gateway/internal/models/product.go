package models

type Product struct {
	Id           string   `json:"id"`
	Name         string   `json:"name,omitempty"`
	Description  string   `json:"description,omitempty"`
	Price        float64  `json:"price,omitempty"`
	Images       []string `json:"images,omitempty"`
	Availability string   `json:"availability,omitempty"`
	Category     string   `json:"category,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	VendorId     string   `json:"vendorId,omitempty"`
}

type StockProduct struct {
	Id           string `json:"id"`
	VendorId     string `json:"vendorId,omitempty"`
	Name         string `json:"name,omitempty"`
	Image        string `json:"image,omitempty"`
	Availability string `json:"availability,omitempty"`
	Quantity     int    `json:"quantity,omitempty"`
}

type IdsToDelete struct {
	Ids []string `json:"ids"`
}

type QuantityUpdate struct {
	Quantity int `json:"quantity"`
}

package models

type Cart struct {
	Id    string `json:"id"`
	Items []Item `json:"items"`
}

type Item struct {
	ProductId string  `json:"productId"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Quantity  int     `json:"quantity"`
}

type CartAddition struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type QuantityUpdate struct {
	Quantity int `json:"quantity"`
}

type CheckoutRequest struct {
	CustomerId      string `json:"customerId"`
	ShippingAddress string `json:"shippingAddress"`
}

type CheckoutResponse struct {
	CheckoutId string `json:"checkoutId"`
	PaymentUrl string `json:"paymentUrl"`
}

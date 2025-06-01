package models

type Profile struct {
	Id              string   `json:"customerId"`
	Email           string   `json:"email"`
	ShippingAddress string   `json:"shippingAddress"`
	Orders          []string `json:"orders"`
}

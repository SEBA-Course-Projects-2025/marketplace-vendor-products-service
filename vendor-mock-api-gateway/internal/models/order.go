package models

import "time"

type Order struct {
	Id               string    `json:"id"`
	CustomerId       string    `json:"customerId"`
	VendorId         string    `json:"vendorId"`
	Items            []string  `json:"items"`
	TotalPrice       float64   `json:"totalPrice"`
	Status           string    `json:"status"`
	Date             time.Time `json:"date"`
	VendorConfStatus string    `json:"vendorConfStatus"`
}

type VendorConfStatusUpdate struct {
	VendorConfStatus string `json:"vendorConfStatus"`
}

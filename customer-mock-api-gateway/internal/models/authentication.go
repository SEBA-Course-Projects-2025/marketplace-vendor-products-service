package models

type Registration struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"shippingAddress"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

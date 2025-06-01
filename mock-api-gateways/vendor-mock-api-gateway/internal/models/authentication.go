package models

type Registration struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Website     string `json:"website,omitempty"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

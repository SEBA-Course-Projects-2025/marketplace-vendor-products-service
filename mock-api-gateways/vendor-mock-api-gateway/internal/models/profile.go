package models

type Profile struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
	Logo        string `json:"logo"`
	Address     string `json:"address"`
	Website     string `json:"website,omitempty"`
	CatalogId   string `json:"catalogId"`
}

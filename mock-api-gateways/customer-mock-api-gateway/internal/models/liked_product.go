package models

type LikedProduct struct {
	ProductId    string  `json:"productId"`
	Name         string  `json:"name"`
	Price        float64 `json:"price"`
	Availability string  `json:"availability"`
	Image        string  `json:"image"`
}

type IdsToDelete struct {
	Ids []string `json:"ids"`
}

type LikedCheck struct {
	Liked bool `json:"liked"`
}

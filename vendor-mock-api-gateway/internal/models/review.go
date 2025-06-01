package models

import "time"

type Review struct {
	Id           string    `json:"id"`
	ProductId    string    `json:"productId"`
	ReviewerId   string    `json:"reviewerId"`
	ReviewerName string    `json:"reviewerName"`
	Rating       float64   `json:"rating"`
	Comment      string    `json:"comment"`
	Date         time.Time `json:"date"`
	Reply        []Reply   `json:"reply"`
}

type Reply struct {
	Id          string    `json:"id"`
	ReplierId   string    `json:"replierId"`
	ReplierName string    `json:"replierName"`
	Comment     string    `json:"comment"`
	Date        time.Time `json:"date"`
}

type ReplyComment struct {
	Comment string `json:"comment"`
}

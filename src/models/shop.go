package models

import "time"

type Shop struct {
	ShopName        string    `json:"shopName"`
	ShopImage       string    `json:"shopImage"`
	ShopDescription string    `json:"shopDescription"`
	ShopRate        uint64    `json:"rate"`
	ShopDate        time.Time `json:"shopDate"`
}

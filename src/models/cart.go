package models

import "time"

type Cart struct {
	CartID   uint      `json:"cartID"`
	Address  string    `json:"address"`
	CartDate time.Time `json:"cartDate"`
	ShopName string    `json:"shopName"`
	Username string    `json:"username"`
	Viewed   bool      `json:"viewed"`
}

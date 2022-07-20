package models

import "time"

type Item struct {
	ItemID          uint      `json:"itemID"`
	ItemName        string    `json:"itemName"`
	ItemImage       string    `json:"itemImage"`
	ItemPrice       uint      `json:"itemPrice"`
	ItemDescription string    `json:"itemDescription"`
	ItemRate        uint64    `json:"rate"`
	ItemDate        time.Time `json:"shopDate"`
	ItemShop        string    `json:"itemShop"`
}

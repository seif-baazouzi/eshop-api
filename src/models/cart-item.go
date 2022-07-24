package models

type CartItem struct {
	CartItemID uint   `json:"cartItemID"`
	Amount     uint   `json:"amount"`
	ItemName   string `json:"itemName"`
	ItemImage  string `json:"itemImage"`
	ItemPrice  uint   `json:"itemPrice"`
}

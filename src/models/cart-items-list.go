package models

type cartItem struct {
	ItemID uint `json:"itemID"`
	Amount uint `json:"amount"`
}

type CartItemsList struct {
	Address  string     `json:"address"`
	ShopName string     `json:"shopName"`
	Items    []cartItem `json:"items"`
}

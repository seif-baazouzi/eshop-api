package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get cart items
// @Success 200 {array} CartItem
// @router /carts/items/:cartID [get]
func GetCartItems(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	cartID, err := url.QueryUnescape(c.Params("cartID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// get cart items list from database
	rows, err := conn.Query(
		"SELECT cartItemID, amount, itemName, itemImage, itemPrice FROM cartsItems C, items I WHERE C.itemID = I.itemID AND C.cartID = $1",
		cartID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	itemsList := []models.CartItem{}
	for rows.Next() {
		var item models.CartItem
		rows.Scan(&item.CartItemID, &item.Amount, &item.ItemName, &item.ItemImage, &item.ItemPrice)
		itemsList = append(itemsList, item)
	}

	return c.JSON(fiber.Map{"cartItems": itemsList})
}

package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get shop carts
// @Success 200 {array} Cart
// @router /carts/:shopName [get]
func GetShopCarts(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	shopName, err := url.QueryUnescape(c.Params("shopName"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// get carts list from database
	rows, err := conn.Query(
		"SELECT cartID, address, cartDate, shopName, username FROM carts WHERE shopName = $1",
		shopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	cartsList := []models.Cart{}
	for rows.Next() {
		var cart models.Cart
		rows.Scan(&cart.CartID, &cart.Address, &cart.CartDate, &cart.ShopName, &cart.Username)
		cartsList = append(cartsList, cart)
	}

	return c.JSON(fiber.Map{"carts": cartsList})
}

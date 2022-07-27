package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get shop carts
// @Success 200 {array} Cart
// @router /carts/:shopName [get]
func GetUserCarts(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")

	// get carts list from database
	rows, err := conn.Query(
		"SELECT cartID, address, cartDate, shopName, username FROM carts WHERE username = $1",
		username,
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

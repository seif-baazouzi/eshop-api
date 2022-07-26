package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get cart items
// @Success 200 {object} message
// @router /carts [post]
func AddCart(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")

	var cart models.CartItemsList

	// parse body
	err := c.BodyParser(&cart)
	if err != nil {
		fmt.Println(err)
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check cart
	errors, err := tests.CheckCart(conn, cart)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if errors != nil {
		return c.JSON(errors)
	}

	// insert cart
	var cartID uint

	err = conn.QueryRow(
		"INSERT INTO carts (address, shopName, username) VALUES ($1, $2, $3) RETURNING cartID",
		cart.Address,
		cart.ShopName,
		username,
	).Scan(&cartID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	// insert items
	for _, item := range cart.Items {
		_, err = conn.Exec(
			"INSERT INTO cartsItems (amount, itemID, cartID) VALUES ($1, $2, $3)",
			item.Amount,
			item.ItemID,
			cartID,
		)

		if err != nil {
			return utils.ServerError(c, err)
		}
	}

	return c.JSON(fiber.Map{"message": "success"})
}

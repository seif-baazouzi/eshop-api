package auth

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

func CheckCartShop(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	cartID, err := url.QueryUnescape(c.Params("cartID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	rows, err := conn.Query(
		"SELECT 1 FROM carts C, shops S WHERE S.owner = $1 AND C.cartID = $2 AND C.shopName = S.shopName",
		username,
		cartID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	if !rows.Next() {
		return c.JSON(fiber.Map{"message": "cart-not-exist"})
	}

	return c.Next()
}

func CheckCartOwner(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	cartID, err := url.QueryUnescape(c.Params("cartID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	rows, err := conn.Query(
		"SELECT 1 FROM carts WHERE username = $1 AND cartID = $2",
		username,
		cartID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	if !rows.Next() {
		return c.JSON(fiber.Map{"message": "cart-not-exist"})
	}

	return c.Next()
}

package auth

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

func CheckShopOwner(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	shopName, err := url.QueryUnescape(c.Params("shopName"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	rows, err := conn.Query(
		"SELECT shopImage FROM shops WHERE shopName = $1 AND owner = $2",
		shopName,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	if !rows.Next() {
		return c.JSON(fiber.Map{"message": "user-not-exist"})
	}

	var shopImage string
	rows.Scan(&shopImage)

	c.Locals("shopImage", shopImage)

	return c.Next()
}

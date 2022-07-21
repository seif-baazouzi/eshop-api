package auth

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

func CheckItemOwner(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	itemID, err := url.QueryUnescape(c.Params("itemID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	rows, err := conn.Query(
		"SELECT shopName FROM shops S, items I WHERE S.owner = $1 AND I.itemID = $2 AND I.shop = S.shopName",
		username,
		itemID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !rows.Next() {
		return c.JSON(fiber.Map{"message": "item-not-exist"})
	}

	var shopName string
	rows.Scan(&shopName)

	c.Locals("shopName", shopName)

	return c.Next()
}

package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get user item rate
// @Success 200 {object} rate
// @router /items/user/rate/:itemID [get]
func GetItemRate(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	itemID, err := url.QueryUnescape(c.Params("itemID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	username := c.Locals("username")

	// get rate
	rows, err := conn.Query(
		"SELECT rate FROM itemsRates WHERE username = $1 AND itemID = $2",
		username,
		itemID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	if rows.Next() {
		var rate uint
		rows.Scan(&rate)

		return c.JSON(fiber.Map{"rate": rate})
	} else {
		return c.JSON(fiber.Map{"rate": 0})
	}
}

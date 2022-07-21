package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description rate an item
// @Success 200 {object} message
// @router /items/:itemID/rate [put]
func ItemRate(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	itemID, err := url.QueryUnescape(c.Params("itemID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	username := c.Locals("username")

	// check rate
	var rate models.Rate
	err = c.BodyParser(&rate)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	if rate.RateValue > 5 {
		return c.JSON(fiber.Map{"message": "invalid-rate-range"})
	}

	// set rate
	rows, err := conn.Query(
		"SELECT 1 FROM itemsRates WHERE username = $1",
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if rows.Next() {
		_, err := conn.Exec(
			"UPDATE itemsRates SET rate = $1 WHERE username = $2",
			rate.RateValue,
			username,
		)

		if err != nil {
			return utils.ServerError(c, err)
		}
	} else {
		_, err := conn.Exec(
			"INSERT INTO itemsRates VALUES ($1, $2, $3)",
			itemID,
			username,
			rate.RateValue,
		)

		if err != nil {
			return utils.ServerError(c, err)
		}
	}

	return c.JSON(fiber.Map{"message": "success"})
}

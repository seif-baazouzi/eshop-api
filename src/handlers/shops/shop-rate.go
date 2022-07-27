package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description rate a shop
// @Success 200 {object} message
// @router /shops/:shopName/rate [put]
func ShopRate(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	shopName, err := url.QueryUnescape(c.Params("shopName"))

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
		"SELECT 1 FROM shopsRates WHERE username = $1",
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	if rows.Next() {
		_, err := conn.Exec(
			"UPDATE shopsRates SET rate = $1 WHERE username = $2",
			rate.RateValue,
			username,
		)

		if err != nil {
			return utils.ServerError(c, err)
		}
	} else {
		_, err := conn.Exec(
			"INSERT INTO shopsRates VALUES ($1, $2, $3)",
			shopName,
			username,
			rate.RateValue,
		)

		if err != nil {
			return utils.ServerError(c, err)
		}
	}

	return c.JSON(fiber.Map{"message": "success"})
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description edit a shop
// @Success 200 {object} message
// @router /shops/:shopName [delete]
func DeleteShop(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	shopName := c.Params("shopName")

	// delete shop
	_, err := conn.Exec(
		"DELETE FROM shops WHERE shopName = $1 AND owner = $2",
		shopName,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

// @Description rate a shop
// @Success 200 {object} message
// @router /shops/:shopName/rate [patch]
func ShopRate(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	shopName := c.Params("shopName")
	username := c.Locals("username")

	// check rate
	var rate models.Rate
	err := c.BodyParser(&rate)

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
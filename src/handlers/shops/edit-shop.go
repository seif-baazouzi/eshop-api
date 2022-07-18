package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description edit a shop
// @Success 200 {object} message
// @router /shops/:shopName [put]
func EditShop(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	shopName := c.Params("shopName")

	// check shop
	var shop models.Shop
	err := c.BodyParser(&shop)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	errors := tests.CheckShop(shop)
	if errors != nil {
		return c.JSON(errors)
	}

	// check if shop is already exist
	rows, err := conn.Query(
		"SELECT 1 FROM shops WHERE shopName = $1",
		shop.ShopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if rows.Next() {
		return c.JSON(fiber.Map{"shopName": "This shop is already exist"})
	}

	// edit shop
	_, err = conn.Exec(
		"UPDATE shops set shopName = $1, shopDescription = $2 WHERE shopName = $3 AND owner = $4",
		shop.ShopName,
		shop.ShopDescription,
		shopName,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description add a shop
// @Success 200 {object} message
// @router /shops [post]
func AddShop(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")

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
	isExist, err := tests.IsShopExists(conn, shop.ShopName)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if isExist {
		return c.JSON(fiber.Map{"shopName": "This shop is already exist"})
	}

	// add shop
	_, err = conn.Exec(
		"INSERT INTO shops (shopName, shopImage, shopDescription, owner) VALUES ($1, '', $2, $3)",
		shop.ShopName,
		shop.ShopDescription,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

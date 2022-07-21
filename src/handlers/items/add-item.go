package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description add an item
// @Success 200 {object} message
// @router /items/:shopName [post]
func AddItem(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	shopName, err := url.QueryUnescape(c.Params("shopName"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check item
	var item models.Item
	err = c.BodyParser(&item)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	errors := tests.CheckItem(item)
	if errors != nil {
		return c.JSON(errors)
	}

	// check if item is already exist
	isExist, err := tests.IsItemExists(conn, item.ItemName, shopName)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if isExist {
		return c.JSON(fiber.Map{"itemName": "This item is already exist"})
	}

	// add shop
	var itemID uint

	err = conn.QueryRow(
		"INSERT INTO items (itemName, itemImage, itemPrice, itemDescription, shop) VALUES ($1, '', $2, $3, $4) RETURNING itemID",
		item.ItemName,
		item.ItemPrice,
		item.ItemDescription,
		shopName,
	).Scan(&itemID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "itemID": itemID})
}

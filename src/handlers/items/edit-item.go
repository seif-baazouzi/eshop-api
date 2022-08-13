package handlers

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description edit an item
// @Success 200 {object} message
// @router /items/:itemID [put]
func EditItem(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	shopName := fmt.Sprintf("%s", c.Locals("shopName"))
	itemName := fmt.Sprintf("%s", c.Locals("itemName"))

	itemID, err := url.QueryUnescape(c.Params("itemID"))

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

	if isExist && item.ItemName != itemName {
		return c.JSON(fiber.Map{"itemName": "This item is already exist"})
	}

	// edit item
	_, err = conn.Exec(
		"UPDATE items SET itemName = $1, itemPrice = $2, itemDescription = $3 WHERE itemID = $4",
		item.ItemName,
		item.ItemPrice,
		item.ItemDescription,
		itemID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description delete item
// @Success 200 {object} message
// @router /items/:itemID [delete]
func DeleteItem(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	itemID, err := url.QueryUnescape(c.Params("itemID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// delete shop
	_, err = conn.Exec(
		"DELETE FROM items WHERE itemID = $1",
		itemID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

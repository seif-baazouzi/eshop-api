package handlers

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description delete shop
// @Success 200 {object} message
// @router /shops/:shopName [delete]
func DeleteShop(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	shopName, err := url.QueryUnescape(c.Params("shopName"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// remove shop image
	oldImageName := fmt.Sprint(c.Locals("shopImage"))
	utils.RemoveImage(oldImageName)

	// delete shop
	_, err = conn.Exec(
		"DELETE FROM shops WHERE shopName = $1",
		shopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

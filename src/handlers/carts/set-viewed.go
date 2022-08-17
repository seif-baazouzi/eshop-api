package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description edit carts viewed status
// @Success 200 {object} message
// @router /carts/:cartID/:status [put]
func SetViewedStatus(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	cartID, err := url.QueryUnescape(c.Params("cartID"))
	status, err := url.QueryUnescape(c.Params("status"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	statusValue := false
	if status == "view" {
		statusValue = true
	}

	// set status
	_, err = conn.Exec(
		"UPDATE carts SET viewed = $1 WHERE cartID = $2",
		statusValue,
		cartID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

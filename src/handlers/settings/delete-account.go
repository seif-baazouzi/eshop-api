package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description delete user account
// @Success 200 {object} message
// @router /users/settings/delete-account [delete]
func DeleteAccount(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")

	_, err := conn.Exec("DELETE FROM users WHERE username = $1", username)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

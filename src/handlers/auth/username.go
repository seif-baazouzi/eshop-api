package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// @Description get username
// @Success 200 {object} username
// @router /username [get]
func GetUsername(c *fiber.Ctx) error {
	username := c.Locals("username")

	return c.JSON(fiber.Map{"username": username})
}

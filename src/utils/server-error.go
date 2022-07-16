package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func ServerError(c *fiber.Ctx, err error) error {
	fmt.Println(err.Error())
	return c.Status(500).JSON(fiber.Map{"message": "server-error"})
}

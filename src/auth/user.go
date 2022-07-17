package auth

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

func IsUserExist(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	email := c.Locals("email")

	rows, err := conn.Query("SELECT username FROM users WHERE email = $1", email)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !rows.Next() {
		return c.JSON(fiber.Map{"message": "user-not-exist"})
	}

	var username string
	rows.Scan(&username)

	c.Locals("username", username)

	return c.Next()
}

package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description update username password
// @Success 200 {object} message
// @router /users/settings/update-username [post]
func UpdateUsername(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	oldUsername := c.Locals("username")

	user := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "invalid-input"})
	}

	errors := tests.CheckUpdateUsername(user.Username, user.Password)

	if errors != nil {
		return c.JSON(errors)
	}

	rows, err := conn.Query("SELECT password as hash FROM users WHERE username = $1", oldUsername)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !rows.Next() {
		return c.JSON(fiber.Map{"username": "This user do not exist"})
	}

	var hash string
	rows.Scan(&hash)

	if !utils.ComparePasswords(user.Password, hash) {
		return c.JSON(fiber.Map{"password": "Wrong password"})
	}

	if user.Username == oldUsername {
		return c.JSON(fiber.Map{"username": "Must not be the same as the old one"})
	}

	rows, err = conn.Query("SELECT password as hash FROM users WHERE username = $1", user.Username)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if rows.Next() {
		return c.JSON(fiber.Map{"username": "This username is already exist"})
	}

	_, err = conn.Exec("UPDATE users SET username = $1 WHERE username = $2", user.Username, oldUsername)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

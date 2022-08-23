package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description update user password
// @Success 200 {object} message
// @router /users/settings/update-password [post]
func UpdatePassword(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")

	payload := struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}{}

	err := c.BodyParser(&payload)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "invalid-input"})
	}

	errors := tests.CheckUpdatePassword(payload.OldPassword, payload.NewPassword)

	if errors != nil {
		return c.JSON(errors)
	}

	rows, err := conn.Query("SELECT password as hash FROM users WHERE username = $1", username)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !rows.Next() {
		return c.JSON(fiber.Map{"username": "This user do not exist"})
	}

	var hash string
	rows.Scan(&hash)

	if !utils.ComparePasswords(payload.OldPassword, hash) {
		return c.JSON(fiber.Map{"oldPassword": "Wrong password"})
	}

	if payload.OldPassword == payload.NewPassword {
		return c.JSON(fiber.Map{"newPassword": "Must not be the same"})
	}

	newPasswordHash := utils.Hash(payload.NewPassword)

	_, err = conn.Exec("UPDATE users SET password = $1 WHERE username = $2", newPasswordHash, username)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

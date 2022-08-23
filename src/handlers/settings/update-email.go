package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description update username email
// @Success 200 {object} message
// @router /users/settings/update-email [post]
func UpdateEmail(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")

	user := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}

	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "invalid-input"})
	}

	errors := tests.CheckUpdateEmail(user.Email, user.Password)

	if errors != nil {
		return c.JSON(errors)
	}

	rows, err := conn.Query("SELECT email, password as hash FROM users WHERE username = $1", username)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !rows.Next() {
		return c.JSON(fiber.Map{"username": "This user do not exist"})
	}

	var oldEmail string
	var hash string
	rows.Scan(&oldEmail, &hash)

	if !utils.ComparePasswords(user.Password, hash) {
		return c.JSON(fiber.Map{"password": "Wrong password"})
	}

	if user.Email == oldEmail {
		return c.JSON(fiber.Map{"email": "Must not be the same as the old one"})
	}

	rows, err = conn.Query("SELECT password as hash FROM users WHERE email = $1", user.Email)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if rows.Next() {
		return c.JSON(fiber.Map{"email": "This email is already exist"})
	}

	_, err = conn.Exec("UPDATE users SET email = $1 WHERE username = $2", user.Email, username)

	if err != nil {
		return utils.ServerError(c, err)
	}

	newToken := auth.GenerateToken("email", user.Email)

	return c.JSON(fiber.Map{"message": "success", "token": newToken})
}

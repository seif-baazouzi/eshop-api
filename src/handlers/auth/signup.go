package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description signup a user
// @Success 200 {object} token
// @router /signup [post]
func UserSignup(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var user models.User

	// parse body
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// validate input
	errors := tests.CheckUserSignup(user)

	if errors != nil {
		return c.JSON(errors)
	}

	// check if the email is exist
	rows, err := conn.Query("SELECT 1 FROM users WHERE email = $1", user.Email)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if rows.Next() {
		return c.JSON(fiber.Map{"email": "This email is already taken"})
	}

	// check if the username is exist
	rows, err = conn.Query("SELECT 1 FROM users WHERE username = $1", user.Username)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	if rows.Next() {
		return c.JSON(fiber.Map{"username": "This username is already taken"})
	}

	// insert user
	hash := utils.Hash(user.Password)

	_, err = conn.Exec(
		"INSERT INTO users VALUES($1, $2, $3)",
		user.Username,
		user.Email,
		hash,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	// response with token
	token := auth.GenerateToken("email", user.Email)
	return c.Status(201).JSON(fiber.Map{"token": token})
}

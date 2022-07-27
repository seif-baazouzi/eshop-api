package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description login a user
// @Success 200 {object} token
// @router /login [post]
func UserLogin(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	var user models.User

	// parse body
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// validate input
	errors := tests.CheckUserLogin(user)

	if errors != nil {
		return c.JSON(errors)
	}

	// check if the user is exist
	rows, err := conn.Query("SELECT password as hash FROM users WHERE email = $1", user.Email)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	if !rows.Next() {
		return c.JSON(fiber.Map{"email": "This user do not exist"})
	}

	// check the password
	var hash string
	rows.Scan(&hash)

	if !utils.ComparePasswords(user.Password, hash) {
		return c.JSON(fiber.Map{"password": "Wrong password"})
	}

	// response with token
	token := auth.GenerateToken("email", user.Email)
	return c.JSON(fiber.Map{"token": token})
}

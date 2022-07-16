package tests

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

func CheckUserLogin(user models.User) fiber.Map {
	errors := make(fiber.Map)

	if user.Email == "" {
		errors["email"] = "Must not be empty"
	} else if !utils.IsValidEmail(user.Email) {
		errors["email"] = "Must be a valid email address"
	}

	if user.Password == "" {
		errors["password"] = "Must not be empty"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}

func CheckUserSignup(user models.User) fiber.Map {
	errors := make(fiber.Map)

	if user.Email == "" {
		errors["email"] = "Must not be empty"
	} else if !utils.IsValidEmail(user.Email) {
		errors["email"] = "Must be a valid email address"
	}

	if user.Username == "" {
		errors["username"] = "Must not be empty"
	} else if !utils.IsValidUsername(user.Username) {
		errors["username"] = "Must be a valid username"
	}

	if user.Password == "" {
		errors["password"] = "Must not be empty"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}

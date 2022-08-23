package tests

import "github.com/gofiber/fiber/v2"

func CheckUpdatePassword(oldPassword string, NewPassword string) fiber.Map {
	errors := make(fiber.Map)

	if oldPassword == "" {
		errors["oldPassword"] = "Must not be empty"
	}

	if NewPassword == "" {
		errors["newPassword"] = "Must not be empty"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}

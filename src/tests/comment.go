package tests

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/models"
)

func CheckComment(comment models.Comment) fiber.Map {
	errors := make(fiber.Map)

	if comment.CommentValue == "" {
		errors["commentValue"] = "Must not be empty"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}

package tests

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/models"
)

func CheckShop(shop models.Shop) fiber.Map {
	errors := make(fiber.Map)

	if shop.ShopName == "" {
		errors["shopName"] = "Must not be empty"
	}

	if shop.ShopDescription == "" {
		errors["shopDescription"] = "Must not be empty"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}

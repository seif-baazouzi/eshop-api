package tests

import (
	"database/sql"

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

func IsShopExists(conn *sql.DB, shopName string) (bool, error) {
	rows, err := conn.Query(
		"SELECT 1 FROM shops WHERE shopName = $1",
		shopName,
	)

	if err != nil {
		return false, err
	}

	defer rows.Close()

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

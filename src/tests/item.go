package tests

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/models"
)

func CheckItem(item models.Item) fiber.Map {
	errors := make(fiber.Map)

	if item.ItemName == "" {
		errors["itemName"] = "Must not be empty"
	}

	if item.ItemPrice == 0 {
		errors["itemPrice"] = "Must be greater than zero"
	}

	if item.ItemDescription == "" {
		errors["itemDescription"] = "Must not be empty"
	}

	if len(errors) != 0 {
		return errors
	}

	return nil
}

func IsItemExists(conn *sql.DB, itemName string, shopName string) (bool, error) {
	rows, err := conn.Query(
		"SELECT 1 FROM items WHERE itemName = $1 AND shop = $2",
		itemName,
		shopName,
	)

	if err != nil {
		return false, err
	}

	if rows.Next() {
		return true, nil
	}

	return false, nil
}

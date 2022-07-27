package tests

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/models"
)

func isExist(arr []uint, item uint) bool {
	for _, el := range arr {
		if el == item {
			return true
		}
	}

	return false
}

func CheckCart(conn *sql.DB, cart models.CartItemsList) (fiber.Map, error) {
	errors := make(fiber.Map)

	// check address
	if cart.Address == "" {
		errors["address"] = "Must not be empty"
	}

	// check shopName
	if cart.ShopName == "" {
		errors["shopName"] = "Must not be empty"
	} else {
		res, err := conn.Query(
			"SELECT 1 FROM shops WHERE shopName = $1",
			cart.ShopName,
		)

		if err != nil {
			return nil, err
		}

		defer res.Close()

		if !res.Next() {
			errors["shopName"] = "Shop does not exist"
		}
	}

	// check Items
	if len(cart.Items) == 0 {
		errors["items"] = "The cart must have one item or more"
	} else {
		itemsIDs := []uint{}
		for _, item := range cart.Items {
			// check itemID
			res, err := conn.Query(
				"SELECT 1 FROM items WHERE itemID = $1 AND shop = $2",
				item.ItemID,
				cart.ShopName,
			)

			if err != nil {
				return nil, err
			}

			defer res.Close()

			if !res.Next() {
				errors["items"] = fmt.Sprintf("Item with id %d does not exist", item.ItemID)
				break
			}

			if isExist(itemsIDs, item.ItemID) {
				errors["items"] = fmt.Sprintf("Item with id %d is duplicated", item.ItemID)
				break
			} else {
				itemsIDs = append(itemsIDs, item.ItemID)
			}

			// check amount
			if item.Amount == 0 {
				errors["items"] = fmt.Sprintf("Amount must be greater than zero in the item with id %d", item.ItemID)
				break
			}
		}
	}

	if len(errors) != 0 {
		return errors, nil
	}

	return nil, nil
}

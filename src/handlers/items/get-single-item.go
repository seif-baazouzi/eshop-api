package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get single item
// @Success 200 {object} Item
// @router /items/:itemID [get]
func GetSingleItems(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	redisClient := db.Redis.Get()
	defer redisClient.Close()

	itemID, err := url.QueryUnescape(c.Params("itemID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// get item from database
	rows, err := conn.Query(
		"SELECT itemID, itemName, itemImage, itemPrice, itemDescription, itemDate, shop FROM items WHERE itemID = $1",
		itemID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	for !rows.Next() {
		return c.JSON(fiber.Map{"message": "Item does not exist"})
	}

	var item models.Item
	rows.Scan(&item.ItemID, &item.ItemName, &item.ItemImage, &item.ItemPrice, &item.ItemDescription, &item.ItemDate, &item.ItemShop)

	rate, err := utils.GetSingleItemRating(conn, redisClient, item.ItemID)

	if err != nil {
		return err
	}

	item.ItemRate = rate

	return c.JSON(fiber.Map{"item": item})
}

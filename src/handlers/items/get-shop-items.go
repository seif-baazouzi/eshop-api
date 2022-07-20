package handlers

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get all items list
// @Success 200 {array} Item
// @router /items/:shopName [get]
func GetShopItems(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	redisClient := db.Redis.Get()
	defer redisClient.Close()

	shopName, err := url.QueryUnescape(c.Params("shopName"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	shopItemsKey := shopName + "ItemsList"

	// check if the itemsList is cached
	res, err := redisClient.Do("GET", shopItemsKey)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if res != nil {
		var result interface{}
		resStr := fmt.Sprintf("%s", res)
		json.Unmarshal([]byte(resStr), &result)

		return c.JSON(fiber.Map{"shopsList": result})
	}

	// get shops list from database
	rows, err := conn.Query(
		"SELECT itemID, itemName, itemImage, itemPrice, itemDescription, itemDate, shop FROM items WHERE shop = $1",
		shopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	itemsList := []models.Item{}
	for rows.Next() {
		var item models.Item
		rows.Scan(&item.ItemID, &item.ItemName, &item.ItemImage, &item.ItemPrice, &item.ItemDescription, &item.ItemDate, &item.ItemShop)
		itemsList = append(itemsList, item)
	}

	// get shops rates
	for index := range itemsList {
		rate, err := utils.GetSingleItemRating(conn, redisClient, itemsList[index].ItemID)

		if err != nil {
			return err
		}

		itemsList[index].ItemRate = rate
	}

	// cache the result
	jsonResult, err := json.Marshal(itemsList)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if len(itemsList) > MIN_ITEMS_COUNT_FOR_CACHING {
		redisClient.Do("SET", shopItemsKey, jsonResult, "EX", "60")
	}

	return c.JSON(fiber.Map{"items": itemsList})
}

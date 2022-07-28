package handlers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get all items list
// @Success 200 {array} Item
// @router /items [get]
func GetAllItems(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	redisClient := db.Redis.Get()
	defer redisClient.Close()

	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil || page < 0 {
		return utils.ServerError(c, err)
	}

	// check if the itemsList is cached
	itemListPageKey := fmt.Sprintf("itemsList-%d", page)

	res, err := redisClient.Do("GET", itemListPageKey)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if res != nil {
		var result interface{}
		resStr := fmt.Sprintf("%s", res)
		json.Unmarshal([]byte(resStr), &result)

		return c.JSON(result)
	}

	// get pages number
	limit := 20
	offset := limit * (page - 1)

	row := conn.QueryRow("SELECT count(*) FROM items")

	var pages int
	row.Scan(&pages)
	pages = pages/limit + 1

	// get items list from database
	rows, err := conn.Query(
		"SELECT itemID, itemName, itemImage, itemPrice, itemDescription, itemDate, shop FROM items ORDER BY itemDate DESC LIMIT $1 OFFSET $2",
		limit,
		offset,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

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
	jsonResult, err := json.Marshal(fiber.Map{"items": itemsList, "pages": pages})

	if err != nil {
		return utils.ServerError(c, err)
	}

	if len(itemsList) != 0 {
		redisClient.Do("SET", itemListPageKey, jsonResult, "EX", "60")
	}

	return c.JSON(fiber.Map{"items": itemsList, "pages": pages})
}

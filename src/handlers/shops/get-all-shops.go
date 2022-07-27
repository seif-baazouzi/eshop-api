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

// @Description get all shops list
// @Success 200 {array} Shop
// @router /shops [get]
func GetAllShops(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	redisClient := db.Redis.Get()
	defer redisClient.Close()

	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil || page < 0 {
		return utils.ServerError(c, err)
	}

	// check if the shopsList is cached
	shopListPageKey := fmt.Sprintf("shopsList-%d", page)

	res, err := redisClient.Do("GET", shopListPageKey)

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

	row := conn.QueryRow("SELECT count(*) FROM shops")

	var pages int
	row.Scan(&pages)
	pages = pages/limit + 1

	// get shops list from database
	rows, err := conn.Query(
		"SELECT shopName, shopImage, shopDescription, shopDate FROM shops ORDER BY shopDate DESC LIMIT $1 OFFSET $2",
		limit,
		offset,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	shopsList := []models.Shop{}
	for rows.Next() {
		var shop models.Shop
		rows.Scan(&shop.ShopName, &shop.ShopImage, &shop.ShopDescription, &shop.ShopDate)
		shopsList = append(shopsList, shop)
	}

	// get shops rates
	err = utils.GetShopsRating(conn, redisClient, shopsList)

	if err != nil {
		return utils.ServerError(c, err)
	}

	// cache the result
	jsonResult, err := json.Marshal(fiber.Map{"shops": shopsList, "pages": pages})

	if err != nil {
		return utils.ServerError(c, err)
	}

	if len(shopsList) != 0 {
		redisClient.Do("SET", shopListPageKey, jsonResult, "EX", "60")
	}

	return c.JSON(fiber.Map{"shops": shopsList, "pages": pages})
}

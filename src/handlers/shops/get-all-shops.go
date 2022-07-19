package handlers

import (
	"encoding/json"
	"fmt"

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

	// check if the shopsList is cached
	res, err := redisClient.Do("GET", "shopsList")

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
	rows, err := conn.Query("SELECT shopName, shopImage, shopDescription, shopDate FROM shops")

	if err != nil {
		return utils.ServerError(c, err)
	}

	shopsList := []models.Shop{}
	for rows.Next() {
		var shop models.Shop
		rows.Scan(&shop.ShopName, &shop.ShopImage, &shop.ShopDescription, &shop.ShopDate)
		shopsList = append(shopsList, shop)
	}

	// get shops rates
	err = utils.GetShopsRating(conn, shopsList)

	if err != nil {
		return utils.ServerError(c, err)
	}

	// cache the result
	jsonResult, err := json.Marshal(shopsList)

	if err != nil {
		return utils.ServerError(c, err)
	}

	redisClient.Do("SET", "shopsList", jsonResult, "EX", "60")

	return c.JSON(fiber.Map{"shops": shopsList})
}

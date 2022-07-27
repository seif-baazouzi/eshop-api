package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get single shop
// @Success 200 {object} Shop
// @router /shops/:shopName [get]
func GetSingleShop(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	redisClient := db.Redis.Get()
	defer redisClient.Close()

	shopName, err := url.QueryUnescape(c.Params("shopName"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// get shop list from database
	rows, err := conn.Query(
		"SELECT shopName, shopImage, shopDescription, shopDate FROM shops WHERE shopName = $1",
		shopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	if !rows.Next() {
		return c.JSON(fiber.Map{"message": "Shop does not exist"})
	}

	var shop models.Shop
	rows.Scan(&shop.ShopName, &shop.ShopImage, &shop.ShopDescription, &shop.ShopDate)

	rate, err := utils.GetSingleShopRating(conn, redisClient, shopName)

	if err != nil {
		return utils.ServerError(c, err)
	}

	shop.ShopRate = rate

	return c.JSON(fiber.Map{"shop": shop})
}

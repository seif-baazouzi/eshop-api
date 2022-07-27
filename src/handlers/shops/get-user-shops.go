package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get user shops list
// @Success 200 {array} Shop
// @router /shops/user [get]
func GetUserShops(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	redisClient := db.Redis.Get()
	defer redisClient.Close()

	username := c.Locals("username")

	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil || page < 0 {
		return utils.ServerError(c, err)
	}

	// get pages number
	limit := 20
	offset := limit * (page - 1)

	row := conn.QueryRow("SELECT count(*) FROM shops WHERE owner = $1", username)

	var pages int
	row.Scan(&pages)
	pages = pages/limit + 1

	// get shops list from database
	rows, err := conn.Query(
		"SELECT shopName, shopImage, shopDescription, shopDate FROM shops WHERE owner = $1 ORDER BY shopDate DESC LIMIT $2 OFFSET $3",
		username,
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

	return c.JSON(fiber.Map{"shops": shopsList, "pages": pages})
}

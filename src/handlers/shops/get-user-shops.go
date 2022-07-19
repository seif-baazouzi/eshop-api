package handlers

import (
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

	username := c.Locals("username")

	// get shops list from database
	rows, err := conn.Query(
		"SELECT shopName, shopImage, shopDescription, shopDate FROM shops WHERE owner = $1",
		username,
	)

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

	return c.JSON(fiber.Map{"shops": shopsList})
}

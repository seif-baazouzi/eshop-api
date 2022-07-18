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
	for index := range shopsList {
		rows, err = conn.Query(
			"SELECT sum(rate) as sum, count(*) as count FROM shops S, shopsRates R WHERE S.shopName = R.shopName AND R.shopName = $1",
			shopsList[index].ShopName,
		)

		if err != nil {
			return utils.ServerError(c, err)
		}

		if rows.Next() {
			var sum uint64
			var count uint64
			rows.Scan(&sum, &count)

			if count == 0 {
				shopsList[index].ShopRate = 0
			} else {
				shopsList[index].ShopRate = sum / count
			}
		}
	}

	return c.JSON(fiber.Map{"shops": shopsList})
}

package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get shop carts
// @Success 200 {array} Cart
// @router /carts/:shopName [get]
func GetUserCarts(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")

	page, err := strconv.Atoi(c.Query("page", "1"))

	if err != nil || page < 0 {
		return utils.ServerError(c, err)
	}

	// get pages number
	limit := 20
	offset := limit * (page - 1)

	row := conn.QueryRow("SELECT count(*) FROM carts WHERE username = $1", username)

	var pages int
	row.Scan(&pages)
	pages = pages/limit + 1

	// get carts list from database
	rows, err := conn.Query(
		"SELECT cartID, address, cartDate, shopName, username FROM carts WHERE username = $1 ORDER BY cartDate DESC LIMIT $2 OFFSET $3",
		username,
		limit,
		offset,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	cartsList := []models.Cart{}
	for rows.Next() {
		var cart models.Cart
		rows.Scan(&cart.CartID, &cart.Address, &cart.CartDate, &cart.ShopName, &cart.Username)
		cartsList = append(cartsList, cart)
	}

	return c.JSON(fiber.Map{"carts": cartsList, "pages": pages})
}

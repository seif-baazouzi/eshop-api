package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"

	"github.com/google/uuid"
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

	// cache the result
	jsonResult, err := json.Marshal(shopsList)

	if err != nil {
		return utils.ServerError(c, err)
	}

	redisClient.Do("SET", "shopsList", jsonResult, "EX", "60")

	return c.JSON(fiber.Map{"shops": shopsList})
}

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

// @Description add a shop
// @Success 200 {object} message
// @router /shops [post]
func AddShop(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")

	// check shop
	var shop models.Shop
	err := c.BodyParser(&shop)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	errors := tests.CheckShop(shop)
	if errors != nil {
		return c.JSON(errors)
	}

	// check if shop is already exist
	rows, err := conn.Query(
		"SELECT 1 FROM shops WHERE shopName = $1",
		shop.ShopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if rows.Next() {
		return c.JSON(fiber.Map{"shopName": "This shop is already exist"})
	}

	// add shop
	_, err = conn.Exec(
		"INSERT INTO shops (shopName, shopImage, shopDescription, owner) VALUES ($1, '', $2, $3)",
		shop.ShopName,
		shop.ShopDescription,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

// @Description edit a shop
// @Success 200 {object} message
// @router /shops/:shopName [put]
func EditShop(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	shopName := c.Params("shopName")

	// check shop
	var shop models.Shop
	err := c.BodyParser(&shop)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	errors := tests.CheckShop(shop)
	if errors != nil {
		return c.JSON(errors)
	}

	// check if shop is already exist
	rows, err := conn.Query(
		"SELECT 1 FROM shops WHERE shopName = $1",
		shop.ShopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if rows.Next() {
		return c.JSON(fiber.Map{"shopName": "This shop is already exist"})
	}

	// edit shop
	_, err = conn.Exec(
		"UPDATE shops set shopName = $1, shopDescription = $2 WHERE shopName = $3 AND owner = $4",
		shop.ShopName,
		shop.ShopDescription,
		shopName,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

// @Description edit a shop
// @Success 200 {object} message
// @router /shops/:shopName [put]
func EditShopImage(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	shopName := c.Params("shopName")

	// check if shop is exist
	rows, err := conn.Query(
		"SELECT shopImage FROM shops WHERE shopName = $1",
		shopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if !rows.Next() {
		return c.JSON(fiber.Map{"shopName": "This shop does not exist"})
	}

	// upload image
	uploadingDir := os.Getenv("UPLOADING_DIRECTORY")

	var oldImageName string
	rows.Scan(&oldImageName)

	if oldImageName != "" {
		os.Remove(fmt.Sprintf("%s/%s", uploadingDir, oldImageName))
	}

	image, err := c.FormFile("image")

	if err != nil {
		return c.JSON(fiber.Map{"message": "Invalid Image"})
	}

	imageName := uuid.New().String() + ".png"

	err = c.SaveFile(image, fmt.Sprintf("%s/%s", uploadingDir, oldImageName))

	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{"message": "Can not upload image"})
	}

	// edit shop image
	_, err = conn.Exec(
		"UPDATE shops set shopImage = $1 WHERE shopName = $2 AND owner = $3",
		imageName,
		shopName,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "image": imageName})
}

// @Description edit a shop
// @Success 200 {object} message
// @router /shops/:shopName [delete]
func DeleteShop(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	shopName := c.Params("shopName")

	// delete shop
	_, err := conn.Exec(
		"DELETE FROM shops WHERE shopName = $1 AND owner = $2",
		shopName,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

// @Description rate a shop
// @Success 200 {object} message
// @router /shops/:shopName/rate [patch]
func ShopRate(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	shopName := c.Params("shopName")
	username := c.Locals("username")

	// check rate
	var rate models.Rate
	err := c.BodyParser(&rate)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	if rate.RateValue > 5 {
		return c.JSON(fiber.Map{"message": "invalid-rate-range"})
	}

	// set rate
	rows, err := conn.Query(
		"SELECT 1 FROM shopsRates WHERE username = $1",
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	if rows.Next() {
		_, err := conn.Exec(
			"UPDATE shopsRates SET rate = $1 WHERE username = $2",
			rate.RateValue,
			username,
		)

		if err != nil {
			return utils.ServerError(c, err)
		}
	} else {
		_, err := conn.Exec(
			"INSERT INTO shopsRates VALUES ($1, $2, $3)",
			shopName,
			username,
			rate.RateValue,
		)

		if err != nil {
			return utils.ServerError(c, err)
		}
	}

	return c.JSON(fiber.Map{"message": "success"})
}

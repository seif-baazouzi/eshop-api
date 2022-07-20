package handlers

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description edit a shop
// @Success 200 {object} message
// @router /shops/:shopName [put]
func EditShopImage(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	shopName, err := url.QueryUnescape(c.Params("shopName"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	oldImageName := fmt.Sprint(c.Locals("shopImage"))

	// upload image
	utils.RemoveImage(oldImageName)

	image, err := c.FormFile("image")

	if err != nil {
		return c.JSON(fiber.Map{"message": "Invalid Image"})
	}

	imageName, err := utils.UploadImage(c, image)

	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{"message": "Can not upload image"})
	}

	// edit shop image
	_, err = conn.Exec(
		"UPDATE shops set shopImage = $1 WHERE shopName = $2",
		imageName,
		shopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "image": imageName})
}

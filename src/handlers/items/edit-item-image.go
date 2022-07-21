package handlers

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description edit item image
// @Success 200 {object} message
// @router /items/:shopName [patch]
func EditItemImage(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	itemID, err := url.QueryUnescape(c.Params("itemID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	oldItemImage := fmt.Sprint(c.Locals("itemImage"))

	// upload image
	utils.RemoveImage(oldItemImage)

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
		"UPDATE items SET itemImage = $1 WHERE itemID = $2",
		imageName,
		itemID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "image": imageName})
}

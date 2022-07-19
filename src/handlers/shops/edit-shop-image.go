package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description edit a shop
// @Success 200 {object} message
// @router /shops/:shopName [put]
func EditShopImage(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	shopName := c.Params("shopName")
	oldImageName := c.Locals("shopImage")

	// upload image
	uploadingDir := os.Getenv("UPLOADING_DIRECTORY")

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
		"UPDATE shops set shopImage = $1 WHERE shopName = $2",
		imageName,
		shopName,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "image": imageName})
}

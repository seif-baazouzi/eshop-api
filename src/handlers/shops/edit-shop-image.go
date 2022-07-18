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

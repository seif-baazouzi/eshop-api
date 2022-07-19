package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func UploadImage(c *fiber.Ctx, image *multipart.FileHeader) (string, error) {
	uploadingDir := os.Getenv("UPLOADING_DIRECTORY")

	imageName := uuid.New().String() + ".png"

	err := c.SaveFile(image, fmt.Sprintf("%s/%s", uploadingDir, imageName))

	if err != nil {
		return "", errors.New("Can not upload image")
	}

	return imageName, nil
}

func RemoveImage(image string) {
	uploadingDir := os.Getenv("UPLOADING_DIRECTORY")

	if image != "" {
		os.Remove(fmt.Sprintf("%s/%s", uploadingDir, image))
	}
}

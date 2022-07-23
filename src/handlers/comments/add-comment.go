package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description add item comment
// @Success 200 {object} message
// @router /comments/:itemID [post]
func AddComments(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	itemID, err := url.QueryUnescape(c.Params("itemID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// check comment
	var comment models.Comment
	err = c.BodyParser(&comment)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	errors := tests.CheckComment(comment)
	if errors != nil {
		return c.JSON(errors)
	}

	// add comment
	var commentID uint

	err = conn.QueryRow(
		"INSERT INTO itemsComments (commentValue, itemID, username) VALUES ($1, $2, $3) RETURNING commentID",
		comment.CommentValue,
		itemID,
		username,
	).Scan(&commentID)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success", "commentID": commentID})
}

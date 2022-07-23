package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/tests"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description edit item comment
// @Success 200 {object} message
// @router /comments/:commentID [put]
func EditComments(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	commentID, err := url.QueryUnescape(c.Params("commentID"))

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

	// edit comment
	_, err = conn.Exec(
		"UPDATE itemsComments SET commentValue = $1 WHERE commentID = $2 AND username = $3",
		comment.CommentValue,
		commentID,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

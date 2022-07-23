package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description delete item comment
// @Success 200 {object} message
// @router /comments/:commentID [delete]
func DeleteComments(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	username := c.Locals("username")
	commentID, err := url.QueryUnescape(c.Params("commentID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// delete comment
	_, err = conn.Exec(
		"DELETE FROM itemsComments WHERE commentID = $1 AND username = $2",
		commentID,
		username,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	return c.JSON(fiber.Map{"message": "success"})
}

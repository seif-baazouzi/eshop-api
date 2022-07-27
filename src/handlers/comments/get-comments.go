package handlers

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/models"
	"gitlab.com/seif-projects/e-shop/api/src/utils"
)

// @Description get item comments
// @Success 200 {array} Comment
// @router /comments/:itemID [get]
func GetComments(c *fiber.Ctx) error {
	conn := db.GetPool()
	defer db.ClosePool(conn)

	itemID, err := url.QueryUnescape(c.Params("itemID"))

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"message": "invalid-input"})
	}

	// get comments list from database
	rows, err := conn.Query(
		"SELECT commentID, commentValue, commentDate, username FROM itemsComments WHERE itemID = $1",
		itemID,
	)

	if err != nil {
		return utils.ServerError(c, err)
	}

	defer rows.Close()

	commentsList := []models.Comment{}
	for rows.Next() {
		var comment models.Comment
		rows.Scan(&comment.CommentID, &comment.CommentValue, &comment.CommentDate, &comment.Username)
		commentsList = append(commentsList, comment)
	}

	return c.JSON(fiber.Map{"comments": commentsList})
}

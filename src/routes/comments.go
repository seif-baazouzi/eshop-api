package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "gitlab.com/seif-projects/e-shop/api/src/handlers/comments"
)

func SetupCommentsRoutes(app *fiber.App) {
	app.Get("/comments/:itemID", handlers.GetComments)
}

package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	handlers "gitlab.com/seif-projects/e-shop/api/src/handlers/comments"
)

func SetupCommentsRoutes(app *fiber.App) {
	app.Get("/comments/:itemID", handlers.GetComments)

	app.Post("/comments/:itemID", auth.IsUser, handlers.AddComments)

	app.Put("/comments/:commentID", auth.IsUser, handlers.EditComments)

	app.Delete("/comments/:commentID", auth.IsUser, handlers.DeleteComments)
}

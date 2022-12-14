package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	handlers "gitlab.com/seif-projects/e-shop/api/src/handlers/auth"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/login", handlers.UserLogin)

	app.Post("/signup", handlers.UserSignup)

	app.Get("/username", auth.IsUser, handlers.GetUsername)
}

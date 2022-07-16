package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/handlers"
)

func SetupAuthRoutes(app *fiber.App) {
	app.Post("/login", handlers.UserLogin)
	app.Post("/signup", handlers.UserSignup)
}

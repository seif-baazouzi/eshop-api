package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	handlers "gitlab.com/seif-projects/e-shop/api/src/handlers/settings"
)

func SetupSettingsRoutes(app *fiber.App) {
	app.Post("/users/settings/update-password", auth.IsUser, handlers.UpdatePassword)

	app.Post("/users/settings/update-username", auth.IsUser, handlers.UpdateUsername)
}

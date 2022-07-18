package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	"gitlab.com/seif-projects/e-shop/api/src/handlers"
)

func SetupShopsRoutes(app *fiber.App) {
	app.Get("/shops", handlers.GetAllShops)
	app.Put("/shops/:shopName/rate", auth.IsUser, handlers.ShopRate)
}

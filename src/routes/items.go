package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "gitlab.com/seif-projects/e-shop/api/src/handlers/items"
)

func SetupItemsRoutes(app *fiber.App) {
	app.Get("/items", handlers.GetAllItems)

	app.Get("/items/shop/:shopName", handlers.GetShopItems)

	app.Get("/items/:itemID", handlers.GetSingleItems)
}

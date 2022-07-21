package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	handlers "gitlab.com/seif-projects/e-shop/api/src/handlers/items"
)

func SetupItemsRoutes(app *fiber.App) {
	app.Get("/items", handlers.GetAllItems)

	app.Get("/items/shop/:shopName", handlers.GetShopItems)

	app.Get("/items/:itemID", handlers.GetSingleItems)

	app.Post("/items/:shopName", auth.IsUser, auth.CheckShopOwner, handlers.AddItem)

	app.Put("/items/:itemID", auth.IsUser, auth.CheckItemOwner, handlers.EditItem)
}

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

	app.Get("/items/user/rate/:itemID", auth.IsUser, handlers.GetItemRate)

	app.Post("/items/:shopName", auth.IsUser, auth.CheckShopOwner, handlers.AddItem)

	app.Put("/items/:itemID", auth.IsUser, auth.CheckItemOwner, handlers.EditItem)

	app.Patch("/items/:itemID", auth.IsUser, auth.CheckItemOwner, handlers.EditItemImage)

	app.Delete("/items/:itemID", auth.IsUser, auth.CheckItemOwner, handlers.DeleteItem)

	app.Put("/items/:itemID/rate", auth.IsUser, handlers.ItemRate)
}

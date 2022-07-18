package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	"gitlab.com/seif-projects/e-shop/api/src/handlers"
)

func SetupShopsRoutes(app *fiber.App) {
	app.Get("/shops", handlers.GetAllShops)
	app.Get("/shops/user", auth.IsUser, handlers.GetUserShops)
	app.Post("/shops", auth.IsUser, handlers.AddShop)
	app.Put("/shops/:shopName", auth.IsUser, handlers.EditShop)
	app.Patch("/shops/:shopName", auth.IsUser, handlers.EditShopImage)
	app.Delete("/shops/:shopName", auth.IsUser, handlers.DeleteShop)
	app.Put("/shops/:shopName/rate", auth.IsUser, handlers.ShopRate)
}

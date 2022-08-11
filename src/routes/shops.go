package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	handlers "gitlab.com/seif-projects/e-shop/api/src/handlers/shops"
)

func SetupShopsRoutes(app *fiber.App) {
	app.Get("/shops", handlers.GetAllShops)

	app.Get("/shops/user", auth.IsUser, handlers.GetUserShops)

	app.Get("/shops/:shopName", handlers.GetSingleShop)

	app.Get("/shops/user/rate/:shopName", auth.IsUser, handlers.GetShopRate)

	app.Post("/shops", auth.IsUser, handlers.AddShop)

	app.Put("/shops/:shopName", auth.IsUser, auth.CheckShopOwner, handlers.EditShop)

	app.Patch("/shops/:shopName", auth.IsUser, auth.CheckShopOwner, handlers.EditShopImage)

	app.Delete("/shops/:shopName", auth.IsUser, auth.CheckShopOwner, handlers.DeleteShop)

	app.Put("/shops/:shopName/rate", auth.IsUser, handlers.ShopRate)
}

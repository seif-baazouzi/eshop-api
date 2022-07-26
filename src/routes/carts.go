package routes

import (
	"github.com/gofiber/fiber/v2"
	"gitlab.com/seif-projects/e-shop/api/src/auth"
	handlers "gitlab.com/seif-projects/e-shop/api/src/handlers/carts"
)

func SetupCartsRoutes(app *fiber.App) {
	app.Get("/carts/user", auth.IsUser, handlers.GetUserCarts)

	app.Get("/carts/:shopName", auth.IsUser, auth.CheckShopOwner, handlers.GetShopCarts)

	app.Get("/carts/items/:cartID", auth.IsUser, auth.CheckCartOwner, handlers.GetCartItems)

	app.Post("/carts", auth.IsUser, handlers.AddCart)
}

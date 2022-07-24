package main

import (
	"os"

	"gitlab.com/seif-projects/e-shop/api/src/db"
	"gitlab.com/seif-projects/e-shop/api/src/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Static("/images", os.Getenv("UPLOADING_DIRECTORY"))

	db.InitPostgresql()
	defer db.ClosePostgresql()

	db.InitRedis()
	defer db.Redis.Close()

	routes.SetupAuthRoutes(app)
	routes.SetupShopsRoutes(app)
	routes.SetupItemsRoutes(app)
	routes.SetupCommentsRoutes(app)
	routes.SetupCartsRoutes(app)

	app.Listen(":3000")
}

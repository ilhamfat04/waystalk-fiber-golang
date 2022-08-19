package main

import (
	"waystalk/database"
	"waystalk/pkg/mysql"
	"waystalk/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	mysql.DatabaseInit()

	database.RunMigration()

	groupRouteApi := app.Group("/api/v1", func(c *fiber.Ctx) error { // middleware for /api/v1
		c.Set("Version", "v1")
		return c.Next()
	})

	routes.RouteInit(groupRouteApi)

	app.Listen("localhost:5000")
}

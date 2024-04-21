package main

import (
	"github.com/dikiwidia/fiber/config"
	"github.com/dikiwidia/fiber/entity/migration"
	"github.com/dikiwidia/fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	config.DatabaseConnect()
	migration.RunMigration()
	app := fiber.New()
	routes.Web(app)
	app.Listen(":3002")
}

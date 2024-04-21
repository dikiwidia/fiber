package main

import (
	"github.com/dikiwidia/fiber/database"
	"github.com/dikiwidia/fiber/database/migration"
	"github.com/dikiwidia/fiber/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.DatabaseConnect()
	migration.RunMigration()
	app := fiber.New()
	routes.Web(app)
	app.Listen(":3002")
}

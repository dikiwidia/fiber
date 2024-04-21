package routes

import (
	"github.com/dikiwidia/fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Web(c *fiber.App) {
	c.Get("/", controllers.HomeIndex)
}

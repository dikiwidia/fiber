package routes

import (
	"github.com/dikiwidia/fiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Web(c *fiber.App) {
	c.Get("/", controllers.HomeControllerIndex)
	c.Get("/users/show", controllers.UserControllerShow)
	c.Post("/users/create", controllers.UserControllerCreate)
	c.Get("/users/data/:id", controllers.UserControllerGetById)
}

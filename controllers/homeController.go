package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func HomeControllerIndex(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}

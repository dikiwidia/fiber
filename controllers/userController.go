package controllers

import (
	"log"

	"github.com/dikiwidia/fiber/database"
	"github.com/dikiwidia/fiber/models/entity"
	"github.com/dikiwidia/fiber/models/req"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func UserControllerIndex(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello World",
	})
}

func UserControllerShow(c *fiber.Ctx) error {
	var users []entity.User
	err := database.DB.Find(&users).Error
	if err != nil {
		log.Println(err)
	}
	return c.JSON(users)
}

func UserControllerCreate(c *fiber.Ctx) error {
	user := new(req.User)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	validation := validator.New()
	if err := validation.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
			"status":  "Failed",
		})
	}

	newUser := entity.User{
		Name:     user.Name,
		Email:    user.Email,
		Age:      user.Age,
		Birthday: user.Birthday,
	}

	if err := database.DB.Create(&newUser).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "Failed",
			"message": "Failed create new user",
		})
	}
	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   newUser,
	})
}

func UserControllerGetById(c *fiber.Ctx) error {
	var user []entity.User
	id := c.Params("id")
	if id == "" {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "ID must be filled",
			"status":  "Failed",
		})
	}

	if err := database.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Not Found",
			"status":  "Failed",
		})
		return nil
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "OK",
		"data":   user,
	})
}

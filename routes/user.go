package routes

import (
	"api/lib"
	"api/models"
	"github.com/gofiber/fiber/v2"
)

type CreateUserBody struct {
	Handle string `validate:"required,min=3,max=30"`
	Name   string `validate:"required,min=3,max=50"`
}

func CreateUserRoute(c *fiber.Ctx) error {
	body := new(CreateUserBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := lib.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	// Making users just as proof of concept
	newUser := models.User{ID: lib.GetNewSnowflake(), Handle: body.Handle, Name: body.Name}
	lib.DB.Create(&newUser)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"msg": "User created",
	})
}

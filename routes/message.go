package routes

import (
	"api/lib"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type SendMessageBody struct {
	// Destination can either be the ID of a user or a group chat
	// It will be interpreted as a proper channel ID later on
	Destination_Id string `validate:"required,number"`
	Content        string `validate:"required,min=1,max=2000"`
}

func SendMessage(c *fiber.Ctx) error {
	body := new(SendMessageBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := lib.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	fmt.Println(body.Destination_Id)
	fmt.Println(body.Content)

	return c.JSON(fiber.Map{
		"ok": "ok",
	})
}

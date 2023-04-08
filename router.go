package main

import (
	"github.com/gofiber/fiber/v2"
	"api/routes"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello world",
		})
	})

	msg_grp := app.Group("/message")
	msg_grp.Post("/send", routes.SendMessage)

}


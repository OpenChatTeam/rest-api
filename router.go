package main

import (
	"api/routes"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "hello world",
		})
	})

	msg_grp := app.Group("/message")
	msg_grp.Post("/send", routes.SendMessageRoute)
	msg_grp.Post("/session", routes.CreateChatSession)

	user_grp := app.Group("/user")
	user_grp.Post("/create", routes.CreateUserRoute)
}

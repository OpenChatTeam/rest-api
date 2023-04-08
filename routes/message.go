package routes

import (
	"api/lib"
	"api/models"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type SendMessageBody struct {
	// Destination can either be the ID of a user or a group chat
	// It will be interpreted as a proper channel ID later on
	Session_Id string `validate:"required,number"`
	Content    string `validate:"required,min=1,max=2000"`
	My_Id      string `validate:"required,number"`
}

func SendMessageRoute(c *fiber.Ctx) error {
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

	// Sending message payload to Redis
	payload, _ := json.Marshal(fiber.Map{
		"senderID":  body.My_Id,
		"sessionID": body.Session_Id,
		"content":   body.Content,
	})
	ctx := context.Background()
	lib.RedisClient.Publish(ctx, "MESSAGE_CREATE", string(payload))

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
}

type CreateChatSessionBody struct {
	My_Id        string `validate:"required,number"`
	Recipient_Id string `validate:"required,number"`
}

func CreateChatSession(c *fiber.Ctx) error {
	body := new(CreateChatSessionBody)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	errors := lib.ValidateStruct(*body)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	sessionId := lib.GetNewSnowflake()
	// Hardcoding for now
	newChatSession := models.ChatSession{ID: sessionId, ChatType: 0, Name: ""}
	lib.DB.Create(&newChatSession)

	// Adding both users to chat session
	var i int64
	i, _ = strconv.ParseInt(body.My_Id, 10, 64)
	lib.DB.Create(&models.UserInChatSessions{UserID: i, ChatSessionID: sessionId})
	i, _ = strconv.ParseInt(body.Recipient_Id, 10, 64)
	lib.DB.Create(&models.UserInChatSessions{UserID: i, ChatSessionID: sessionId})

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{})
}

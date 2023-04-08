package main

import (
	"fmt"
	"math/rand"
	"os"

	"api/lib"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitApp() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
	})

	app.Use(cors.New())

	return app
}

func main() {
	app := InitApp()

	lib.InitDatabase()
	lib.InitSnowflakeNode(rand.Int63())
	lib.InitRedisClient()
	RegisterRoutes(app)

	app.Listen(fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT")))
}

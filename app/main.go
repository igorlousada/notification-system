package main

import (
    "log"

    "github.com/gofiber/fiber/v3"
	kafka "notification-system/kafka"
)

func main() {
    app := fiber.New()

	app.Post("/send-notification", func(c fiber.Ctx) error {
		kafka.PublishTopic("test-topic", "Hello Moto!")
		return c.SendString("abacate")
	})

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}
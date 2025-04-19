package main

import (
    "log"
	"os"
    "github.com/gofiber/fiber/v3"
	kafka "notification-system/kafka"
)

type publishInterface interface {
	PublishTopic(topic string, message string) (string, error)
}

type messageService struct {
	service publishInterface
}

type mockType struct {}

func (c *mockType) PublishTopic(topic string, message string) (string, error ) {
	return "ok", nil
}


type kafkaType struct {}

func (c *kafkaType) PublishTopic(topic string, message string) (string, error ) {
	ok, err := kafka.PublishTopic(topic, message)
	if err != nil {
		return "", err
	}

	return ok, nil
}

func NewService(m publishInterface) *messageService {
	return &messageService{service: m}
}

func main() {
   app := setupRoutes()
   log.Fatal(app.Listen(":3000"))
}

func setupRoutes() *fiber.App {
	app := fiber.New()

	app.Post("/send-notification", func(c fiber.Ctx) error {
		if os.Getenv("test_env") == "true" {
			mock := &mockType{}
			mockMessage := NewService(mock)
			mockMessage.service.PublishTopic("test-topic", "Hello kafka!")
		} else {
			test := &kafkaType{}
			testMessage := NewService(test)
			testMessage.service.PublishTopic("test-topic", "Hello kafka!")
		}
		return c.SendString("abacate")
	})

    return app
}
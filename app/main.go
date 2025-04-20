package main

import (
    "log"
	"os"
	"encoding/json"

	kafka "notification-system/kafka"
	"github.com/gofiber/fiber/v3"
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
		var msg Message

		log.Printf("%+v", string(c.Body()))
		err := json.Unmarshal(c.Body(), &msg)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "bad request",
				"message": "JSON body is malformed",
			})
		}

		if msg.User_uuid == "" {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "invalid request",
				"message": "User_uuid is empty",
			})
		}

		if msg.Message == "" {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"error": "invalid request",
				"message": "Message is empty",
			})
		}

		var publisher publishInterface

		if os.Getenv("test_env") == "true" {
			publisher = &mockType{}
		} else {
			publisher = &kafkaType{}
		}

		sendMessage := NewService(publisher)

		ok, err := sendMessage.service.PublishTopic("notification-purchase", string(c.Body()))

		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
				"message": "Server could not process the message",
			})
		}

		if ok != "ok" {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
				"message": "Server could not process the message",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "OK",
			"message": "Notification sent",
		})
	})

    return app
}


type Message struct {
	User_uuid string `json: "user_uuid"`
	Message string	`json: "message"`
}
package main

import (
    "log"
	"os"
    "github.com/gofiber/fiber/v3"
	// "encoding/json"
	// "fmt"
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
		// var msg Message

		// err := json.Unmarshal(c.Body(), &msg)

		// if err != nil {
		// 	return c.SendString("invalid request")
		// }

		// fmt.Printf("%+v\n", msg)
		// fmt.Printf(string(c.Body()))
		if os.Getenv("test_env") == "true" {
			mock := &mockType{}
			mockMessage := NewService(mock)
			mockMessage.service.PublishTopic("notification-purchase", string(c.Body()))
		} else {
			test := &kafkaType{}
			testMessage := NewService(test)
			testMessage.service.PublishTopic("notification-purchase", string(c.Body()))
		}
		return c.SendString("abacate")
	})

    return app
}


// type Message struct {
// 	User_uuid string `json: "user_uuid"`
// 	Message string	`json: "message"`
// }
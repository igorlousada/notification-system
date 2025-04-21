package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	kafka "notification-system/kafka"
	"github.com/IBM/sarama"
)

type Consumer struct{}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var err error
		log.Printf("Message claimed: topic = %s, partition = %d, offset = %d, value = %s",
			message.Topic, message.Partition, message.Offset, string(message.Value))
		if os.Args[2] == "fanout" {
			_, err = kafka.PublishTopic("email-topic", string(message.Value))
		} else {
			_, err = sendEmail(string(message.Value))
		}

		if err != nil {
			log.Printf("could not process worker %s: %s", os.Args[2], os.Args[1])
			return nil
		}
			
		session.MarkMessage(message, "")
		session.Commit()
	}
	return nil
}

func sendEmail(message string) (string, error){
	log.Printf("sending email - %s", message)
	return "ok", nil
}

func main() {
	
	topic := os.Args[1]

	client, err := kafka.BuildClient()
	
	consumer := &Consumer{}

	ctx, cancel := context.WithCancel(context.Background())
	
	if err != nil {
		log.Fatalf("Error creating consumer group client: %v", err)
	}
	defer client.Close()

	go func() {
		for {
			if err := client.Consume(ctx, []string{topic}, consumer); err != nil {
				log.Fatalf("Error during consuming: %v", err)
			}
		}
	}()

	log.Println("Consumer started. Press Ctrl+C to stop.")
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, os.Interrupt)
	<-sigterm
	cancel()
	log.Println("Consumer stopped.")
}
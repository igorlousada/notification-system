package Kafka

import (
	"context"
	"log"
	"os"
	"os/signal"
	// "fmt"

	"github.com/IBM/sarama"
)

type Consumer struct{}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

// This is called once per message
func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: topic = %s, partition = %d, offset = %d, value = %s",
			message.Topic, message.Partition, message.Offset, string(message.Value))
			PublishTopic("email-topic", string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}

func main() {
	brokers := []string{"localhost:9092"}
	topic := "test-topic"
	group := "your-consumer-group"

	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0 // pick appropriate version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer := &Consumer{}

	ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(brokers, group, config)
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
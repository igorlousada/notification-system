package Kafka

import (
	"github.com/IBM/sarama"
)



func BuildClient() (sarama.ConsumerGroup, error) {
	brokers := []string{"localhost:9092"}
	// topic := "test-topic"
	group := "your-consumer-group"

	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0 // pick appropriate version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	// consumer := &Consumer{}

	// ctx, cancel := context.WithCancel(context.Background())
	client, err := sarama.NewConsumerGroup(brokers, group, config)
	
	if err != nil {
		return nil, err
	}

	return client, nil
}
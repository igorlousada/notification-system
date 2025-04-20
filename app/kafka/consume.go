package Kafka

import (
	"github.com/IBM/sarama"
)



func BuildClient() (sarama.ConsumerGroup, error) {
	brokers := []string{"localhost:9092"}
	group := "your-consumer-group"

	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetNewest
	config.Consumer.Offsets.AutoCommit.Enable = false

	client, err := sarama.NewConsumerGroup(brokers, group, config)
	
	if err != nil {
		return nil, err
	}

	return client, nil
}
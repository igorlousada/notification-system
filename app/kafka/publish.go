package Kafka
	
	import (
		"log"
		"github.com/IBM/sarama"
	)

	func PublishTopic(topic string, messageToSend string) (string, error) {
		config := sarama.NewConfig()
		config.Producer.RequiredAcks = sarama.WaitForAll
		config.Producer.Return.Successes = true
		config.Producer.Retry.Max = 5
		producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, config)
		if err != nil {
			log.Fatalf("Error while creating producer: %v", err)
		}

		defer producer.Close()

		message := &sarama.ProducerMessage {
			Topic: topic,
			Value: sarama.StringEncoder(messageToSend),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			log.Fatalf("Error producer: %v", err)
		}

		log.Printf("Message sent to partition %d at offset %d\n", partition, offset)

		return "ok", nil
	}
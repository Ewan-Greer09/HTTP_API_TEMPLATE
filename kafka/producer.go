package kafka

import (
	"log"

	"github.com/Shopify/sarama"
)

type Producer struct{}

func (p *Producer) PushCommentToQueue(topic string, message []byte) error {
	brokersURL := []string{"localhost:9092"}
	producer, err := ConnectProducer(brokersURL)
	if err != nil {
		return err
	}

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)", topic, partition, offset)

	return nil
}

func ConnectProducer(brokersURL []string) (sarama.SyncProducer, error) {
	// connect to kafka
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokersURL, config)
	if err != nil {
		return nil, err
	}

	return producer, nil
}

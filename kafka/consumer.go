package kafka

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

type Consumer struct{}

// ConsumeMessage is a function to consume message from kafka
func (c *Consumer) ConsumeMessage() {
	// connect to kafka
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()

	// get partition
	partitionList, err := consumer.Partitions("test")
	if err != nil {
		panic(err)
	}

	// consume message
	for partition := range partitionList {
		pc, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetNewest)
		if err != nil {
			panic(err)
		}
		defer pc.AsyncClose()

		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				fmt.Println(string(msg.Value))
			}
		}(pc)
	}

	// wait for a while
	time.Sleep(10 * time.Second)
}

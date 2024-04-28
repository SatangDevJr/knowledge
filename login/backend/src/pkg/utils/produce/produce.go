package produce

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Produce interface {
	Producer() error
}

func Producer(messages []string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "10.116.1.79:19093"})
	if err != nil {
		fmt.Println("Producer Error :", err)
	}

	defer p.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	topic := "ptp-first-so-hot"
	for _, word := range messages {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)
	return nil
}

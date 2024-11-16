package consumer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"go-kafka-io/model"
	"log"
)

const KafkaNetwork = "tcp"

// EventConsumer An abstract interface declaring operations for consuming messages
type EventConsumer interface {
	ConsumeAll(ctx context.Context)
}

type KafkaEventConsumer struct {
	Address   string
	Topic     string
	Partition int
}

func NewKafkaEventConsumer(Address string, Topic string, Partition int) *KafkaEventConsumer {
	return &KafkaEventConsumer{
		Address:   Address,
		Topic:     Topic,
		Partition: Partition,
	}
}

func (kc *KafkaEventConsumer) ConsumeAll(ctx context.Context) {

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{kc.Address},
		Topic:     kc.Topic,
		Partition: kc.Partition,
	})

	defer func(kafkaReader *kafka.Reader) {
		err := kafkaReader.Close()
		if err != nil {
			log.Println("Error closing kafka reader", err)
		}
	}(kafkaReader)

	for {
		kafkaEvent, err := kafkaReader.ReadMessage(ctx)
		if &kafkaEvent != nil && kafkaEvent.Value == nil {
			break
		}
		if err != nil {
			log.Fatalln("Error reading kafka event:", err)
		}

		var genericEvent model.GenericEvent

		// Expensive - demo purposes only
		err = json.Unmarshal(kafkaEvent.Value, &genericEvent)
		if err != nil {
			log.Fatalln("Failed to unmarshal event:", err)
		}

		if err := kafkaReader.Close(); err != nil {
			log.Fatalln("Failed to close Kafka reader:", err)
		}

		log.Println("Event received:", genericEvent)
	}
}

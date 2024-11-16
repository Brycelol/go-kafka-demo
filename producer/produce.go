package producer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"go-kafka-io/model"
	"log"
)

const KafkaNetwork = "tcp"

// EventProducer An abstract interface declaring operations for producing messages
type EventProducer interface {
	ProduceEvent(event model.GenericEvent) error
}

type KafkaEventProducer struct {
	Address   string
	Topic     string
	Partition int
}

func NewKafkaEventProducer(Address string, Topic string, Partition int) *KafkaEventProducer {
	return &KafkaEventProducer{
		Address:   Address,
		Topic:     Topic,
		Partition: Partition,
	}
}

func (kp *KafkaEventProducer) ProduceEvent(ctx context.Context, event *model.GenericEvent) error {
	log.Println("Producing event: ", event)

	conn, err := kafka.DialLeader(ctx, KafkaNetwork, kp.Address, kp.Topic, kp.Partition)
	if err != nil {
		log.Println("Failed to dial leader:", err)
		return err
	}

	defer func(conn *kafka.Conn) {
		err := conn.Close()
		if err != nil {
			log.Println("Failed to close leader connection:", err)
		}
	}(conn)

	// Expensive - use JSON for demonstration purposes.
	eventJsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Println("Failed to marshal event:", err)
		return err
	}

	_, err = conn.WriteMessages(kafka.Message{
		Value: eventJsonBytes,
	})
	if err != nil {
		log.Println("Failed to write message", err)
		return err
	}

	return nil
}

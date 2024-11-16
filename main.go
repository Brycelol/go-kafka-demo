package main

import (
	"context"
	"github.com/google/uuid"
	"go-kafka-io/consumer"
	"go-kafka-io/model"
	"go-kafka-io/producer"
	"log"
	"sync"
)

// EventCount declares how many events our simulator will publish and consume.
const EventCount = 5
const KafkaAddress = "localhost:19092"
const KafkaTopic = "gareth-events"
const KafkaPartition = 0

func main() {

	ctx := context.Background()

	kafkaProducer := producer.NewKafkaEventProducer(KafkaAddress, KafkaTopic, KafkaPartition)

	var publishWaitGroup sync.WaitGroup

	publishWaitGroup.Add(EventCount)

	// Produce n events asynchronously
	for i := 0; i < EventCount; i++ {
		go func() {
			defer publishWaitGroup.Done()

			message := uuid.NewString() // Junk data to populate message
			mockEvent := model.NewGenericEvent(int64(i+1), message)

			err := kafkaProducer.ProduceEvent(ctx, mockEvent)
			if err != nil {
				log.Fatalf("Failed to produce event: %v", err)
			}
		}()
	}

	publishWaitGroup.Wait()

	// Consume all events synchronously for demonstration purposes.
	kafkaConsumer := consumer.NewKafkaEventConsumer(KafkaAddress, KafkaTopic, KafkaPartition)
	// Drains topic and dumps to STDOUT
	kafkaConsumer.ConsumeAll(ctx)
}

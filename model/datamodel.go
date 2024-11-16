package model

import (
	"fmt"
	"time"
)

// GenericEvent struct representing data we want to shift around Kafka topics
type GenericEvent struct {
	ID        int64
	Message   string
	timestamp time.Time
}

// NewGenericEvent Returns a new Instance of a Generic Event based on specified parameters
func NewGenericEvent(ID int64, Message string) *GenericEvent {
	return &GenericEvent{
		ID:        ID,
		Message:   Message,
		timestamp: time.Now(),
	}
}

// String Stringifies the Generic Event
func (g GenericEvent) String() string {
	return fmt.Sprintf("GenericEvent[ID=%d, Message=%s, timestamp=%s]", g.ID, g.Message, g.timestamp)
}

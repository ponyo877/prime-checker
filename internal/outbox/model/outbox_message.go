package model

import (
	"encoding/json"
	"time"
)

type OutboxMessage struct {
	id        int32
	eventType string
	payload   json.RawMessage
	processed bool
	createdAt time.Time
	updatedAt time.Time
}

func NewOutboxMessage(id int32, eventType string, payload json.RawMessage, processed bool, createdAt, updatedAt time.Time) *OutboxMessage {
	return &OutboxMessage{
		id:        id,
		eventType: eventType,
		payload:   payload,
		processed: processed,
		createdAt: createdAt,
		updatedAt: updatedAt,
	}
}

func (o *OutboxMessage) ID() int32 {
	return o.id
}

func (o *OutboxMessage) EventType() string {
	return o.eventType
}

func (o *OutboxMessage) Payload() json.RawMessage {
	return o.payload
}

func (o *OutboxMessage) IsProcessed() bool {
	return o.processed
}

func (o *OutboxMessage) CreatedAt() time.Time {
	return o.createdAt
}

func (o *OutboxMessage) UpdatedAt() time.Time {
	return o.updatedAt
}

func (o *OutboxMessage) MarkAsProcessed(now time.Time) {
	o.processed = true
	o.updatedAt = now
}

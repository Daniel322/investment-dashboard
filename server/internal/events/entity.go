package events

import (
	"time"

	"github.com/google/uuid"
)

type Event[T any] struct {
	ID        int64      `json:"id"`
	Event     string     `json:"event"`
	Type      string     `json:"type"`
	PayloadId *uuid.UUID `json:"payload_id"`
	CreatedAt time.Time  `json:"created_at"`
	Data      T          `json:"data"`
}

type EventRecord struct {
	ID        int64       `json:"id"`
	Event     string      `json:"event"`
	Type      string      `json:"type"`
	PayloadId string      `json:"payload_id"`
	CreatedAt time.Time   `json:"created_at"`
	Data      interface{} `json:"data"`
}

func Create[T any](event string, type_ string, payloadId *uuid.UUID, data T) (*Event[T], error) {
	eventEntity, err := constructor(0, event, type_, payloadId, time.Now(), data)
	if err != nil {
		return nil, err
	}
	return eventEntity, nil
}

func (e *Event[T]) ToRecord() *EventRecord {
	return &EventRecord{
		ID:        e.ID,
		Event:     e.Event,
		Type:      e.Type,
		PayloadId: e.PayloadId.String(),
		CreatedAt: e.CreatedAt,
		Data:      e.Data,
	}
}

func constructor[T any](
	id int64,
	event string,
	type_ string,
	payloadId *uuid.UUID,
	createdAt time.Time,
	data T,
) (*Event[T], error) {
	return &Event[T]{
		ID:        id,
		Event:     event,
		Type:      type_,
		PayloadId: payloadId,
		CreatedAt: createdAt,
		Data:      data,
	}, nil
}

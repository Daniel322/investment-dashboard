package interfaces

import "fmt"

type Event[T any] struct {
	Topic   string
	Payload T
}

type EventHandler[T any] interface {
	Handle(event Event[T]) error
}

type eventHandlerFunc func(event Event[any]) error

func (f eventHandlerFunc) Handle(event Event[any]) error {
	return f(event)
}

func AdaptHandler[T any](handler EventHandler[T]) EventHandler[any] {
	return eventHandlerFunc(func(event Event[any]) error {
		payload, ok := event.Payload.(T)
		if !ok {
			return fmt.Errorf("unexpected payload type for %s: %T", event.Topic, event.Payload)
		}

		return handler.Handle(Event[T]{
			Topic:   event.Topic,
			Payload: payload,
		})
	})
}

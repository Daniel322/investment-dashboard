package interfaces

type Event[T any] struct {
	Topic   string
	Payload T
}

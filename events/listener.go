package events

type EventListener interface {
	Listen(exchange string, eventNames ...string) (<-chan Event, <-chan error, error)
}

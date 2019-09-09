package amqp

import (
	"encoding/json"

	"github.com/streadway/amqp"
	"github.com/waleplanet/alpha_shared/events"
)

type EventEmitter interface {
	Emit(event events.Event, exchange string) error
}

// package-private
type amqpEventEmitter struct {
	connection *amqp.Connection
}

func (a *amqpEventEmitter) setup(exchange string) error {
	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	return channel.ExchangeDeclare(exchange, "topic", true, false, false, false, nil)

}

func NewAMQPEventEmitter(conn *amqp.Connection, exchange string) (EventEmitter, error) {
	emitter := &amqpEventEmitter{
		connection: conn,
	}

	err := emitter.setup(exchange)
	if err != nil {
		return nil, err
	}
	return emitter, nil
}

func (a *amqpEventEmitter) Emit(event events.Event, exchange string) error {
	jsonData, err := json.Marshal(event)
	if err != nil {
		return err
	}

	channel, err := a.connection.Channel()
	if err != nil {
		return err
	}
	msg := amqp.Publishing{
		Headers: amqp.Table{
			"x-event-name": event.EventName(),
		},
		Body:        jsonData,
		ContentType: "application/json",
	}

	return channel.Publish(exchange, event.EventName(), true, true, msg)

}

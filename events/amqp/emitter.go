package amqp

import "github.com/streadway/amqp"

type EventEmitter interface {
	Emit(event Event) error
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

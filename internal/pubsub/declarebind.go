package pubsub

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

type SimpleQueueType string

const (
	DurableQueue   SimpleQueueType = "durable"
	TransientQueue SimpleQueueType = "transient"
)

func DeclareAndBind(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	queueType SimpleQueueType, // SimpleQueueType is an "enum" type I made to represent "durable" or "transient"
) (*amqp.Channel, amqp.Queue, error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	var durable bool
	switch queueType {
	case DurableQueue:
		durable = true
	case TransientQueue:
		durable = false
	default:
		return nil, amqp.Queue{}, fmt.Errorf("invalid queue type: %s", queueType)
	}

	q, err := ch.QueueDeclare(
		queueName,
		durable,
		queueType == TransientQueue, // autoDelete is true for transient queues
		queueType == TransientQueue, // exclusive is true for transient queues
		false,
		nil,
	)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	err = ch.QueueBind(
		q.Name,
		key,
		exchange,
		false,
		nil,
	)
	if err != nil {
		return nil, amqp.Queue{}, err
	}

	return ch, q, nil
}

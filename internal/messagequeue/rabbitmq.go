package messagequeue

import (
	"chat-chat-go/internal/models"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	connToPublish *amqp.Connection
	connToConsume *amqp.Connection
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	connToPublish, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	connToConsume, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		connToPublish: connToPublish,
		connToConsume: connToConsume,
	}, nil
}

func (r *RabbitMQ) Publish(queueName string, message models.QueueMessage) error {
	rch, err := r.connToPublish.Channel()
	if err != nil {
		return err
	}
	defer rch.Close()

	_, err = rch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgBytes, _ := json.Marshal(message)
	rch.Publish("", queueName, true, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msgBytes,
	})

	return nil
}

func (r *RabbitMQ) Consume(queueName string) (<-chan models.QueueMessage, error) {
	// noWait ??
	// why separate connections for pub and cons ??
	rch, err := r.connToConsume.Channel()
	if err != nil {
		// log.Fatalf("Failed to open channel to consume: %v", err)
		return nil, err
	}

	_, err = rch.QueueDeclare(queueName, true, true, false, false, nil)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	deliveries, err := rch.Consume(queueName, "", true, true, false, false, nil)
	if err != nil {
		// log.Fatalf("Failed to consume from queue: %s : %v", queueName, err)
		return nil, err
	}

	ch := make(chan models.QueueMessage)

	go func() {
		for d := range deliveries {
			var objMsg models.QueueMessage
			err = json.Unmarshal(d.Body, &objMsg)
			if err != nil {
				log.Fatalf("Failed to deserialize the message: %v", err)
			}
			ch <- objMsg
		}
	}()

	return ch, nil
}

func (r *RabbitMQ) Close() {
	r.connToPublish.Close()
	r.connToConsume.Close()
}

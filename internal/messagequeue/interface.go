package messagequeue

import "chat-chat-go/internal/models"

type MessageQueue interface {
	Publish(queueName string, message models.QueueMessage) error
	Consume(queue string) (<-chan models.QueueMessage, error)
	Close()
}

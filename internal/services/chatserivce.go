package services

import (
	"chat-chat-go/internal/messagequeue"
	userregistry "chat-chat-go/internal/repository/userregistry"
	"log"
	"os"
)

var instanceId string
var rabbitMQ messagequeue.MessageQueue
var userRegistry userregistry.UserRegistry

func connectToMessageQueue() {
	var err error
	rabbitMQ, err = messagequeue.NewRabbitMQ("amqp://127.0.0.1")
	if err != nil {
		log.Fatal("Can not connect message queue:", err)
	}
}

func connectToUserRegistry() {
	userRegistry = userregistry.NewRedisRepository("localhost:6379")
}

func InitService() {
	connectToMessageQueue()
	connectToUserRegistry()
	instanceId = os.Getenv("INSTANCE_ID")

	ch, err := rabbitMQ.Consume(InstanceIdToQueueName(instanceId))

	if err != nil {
		log.Fatalln("Failed to consume:", err)
		return
	} else {
		log.Printf("Instace %s now has a queue\n", instanceId)
	}

	go func() {
		for msg := range ch {
			log.Printf("%s is saying to %s : \"%s\"", msg.FromUser, msg.ToUser, msg.Content)
		}
	}()

}

func InstanceIdToQueueName(instanceId string) string {
	return instanceId
}

func SetUserRegistry(userId string) {
}

func HandleUserConnection(userId string) {
	userRegistry.RegisterUser(userId, InstanceIdToQueueName(instanceId))
}

func HandleUserDisconnection(userId string) {
	userRegistry.UnregisterUser(userId)
}

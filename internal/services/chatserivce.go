package services

import (
	"chat-chat-go/internal/messagequeue"
	"chat-chat-go/internal/models"
	"chat-chat-go/internal/repository/chatrepo"
	userregistry "chat-chat-go/internal/repository/userregistry"
	"chat-chat-go/internal/services/userservice"
	"chat-chat-go/internal/websocketmanager"
	"log"
	"os"
)

type ChatService struct {
	instanceId       string
	messageQueue     messagequeue.MessageQueue
	userRegistry     userregistry.UserRegistry
	webSocketManager *websocketmanager.WebSocketManager
	chatRepository   chatrepo.ChatRepository
	userService      *userservice.UserService
}

func NewChatService(webSocketManager *websocketmanager.WebSocketManager) *ChatService {
	return &ChatService{
		webSocketManager: webSocketManager,
		userService:      userservice.NewUserService(),
	}
}

func (chatService *ChatService) connectToMessageQueue() {
	var err error
	chatService.messageQueue, err = messagequeue.NewRabbitMQ(os.Getenv("MESSAGE_QUEUE_URL"))
	if err != nil {
		log.Fatal("Can not connect message queue:", err)
	}
}

func (chatService *ChatService) connectToChatRepository() {
	chatService.chatRepository = chatrepo.NewScyllaChatRepository(os.Getenv("CHAT_REPO_URL"))
}

func (chatService *ChatService) connectToUserRegistry() {
	chatService.userRegistry = userregistry.NewRedisRepository(os.Getenv("USER_REGISTRY_URL"))
}

func (chatService *ChatService) InitService() {
	chatService.connectToMessageQueue()
	chatService.connectToUserRegistry()
	chatService.connectToChatRepository()
	chatService.instanceId = os.Getenv("INSTANCE_ID")

	ch, err := chatService.messageQueue.Consume(chatService.InstanceIdToQueueName(chatService.instanceId))

	if err != nil {
		log.Fatalln("Failed to consume:", err)
		return
	} else {
		log.Printf("Instace %s now has a queue\n", chatService.instanceId)
	}

	go func() {
		for msg := range ch {
			chatService.HandleReceiveMessage(msg)
		}
	}()
}

func (chatService *ChatService) InstanceIdToQueueName(instanceId string) string {
	return instanceId
}

func (chatService *ChatService) HandleUserConnection(userId string) {
	if err := chatService.userRegistry.RegisterUser(userId, chatService.InstanceIdToQueueName(chatService.instanceId)); err != nil {
		log.Println("Failed to register user: ", err)
	}
}

func (chatService *ChatService) HandleUserDisconnection(userId string) {
	chatService.userRegistry.UnregisterUser(userId)
}

func (chatService *ChatService) HandleSendMessage(userId string, message models.DtoInChatSocketMessage) {
	queueName, err := chatService.userRegistry.GetUserRegistry(message.ToUser)
	if err != nil {
		if chatService.userRegistry.IsNonExistingError(err) {
			log.Println("User not registered")
		} else {
			log.Println("Failed to find user registry: ", err)
		}
		return
	}

	err = chatService.messageQueue.Publish(queueName, models.QueueMessage{
		FromUser: userId,
		ToUser:   message.ToUser,
		Content:  message.Content,
	})

	if err != nil {
		log.Println("Failed to send message: ", err)
	}
}

func (chatService *ChatService) HandleReceiveMessage(message models.QueueMessage) {
	chatService.webSocketManager.SendMessage(message.ToUser, models.DtoOutChatSocketMessage(message))
}

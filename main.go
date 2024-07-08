package main

import (
	"chat-chat-go/internal/handlers"
	router "chat-chat-go/internal/router"
	service "chat-chat-go/internal/services"
	"chat-chat-go/internal/websocketmanager"
	"os"

	"github.com/google/uuid"
)

func main() {
	os.Setenv("INSTANCE_ID", uuid.NewString())
	socketManager := websocketmanager.NewWebSocketManager()
	chatService := service.NewChatService(socketManager)
	socketHandler := handlers.NewWebSocketHandler(chatService, socketManager)

	chatService.InitService()
	router.InitRouter(*socketHandler)
}

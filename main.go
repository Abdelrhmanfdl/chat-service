package main

import (
	router "chat-chat-go/internal/router"
	service "chat-chat-go/internal/services"
	"os"

	"github.com/google/uuid"
)

func main() {
	os.Setenv("INSTANCE_ID", uuid.NewString())
	service.InitService()
	router.InitRouter()
}

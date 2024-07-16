package router

import (
	handler "chat-chat-go/internal/handlers"
	middlewares "chat-chat-go/internal/middlewares"
	"os"

	"github.com/gin-gonic/gin"
)

func InitRouter(webSocketHandler *handler.WebSocketHandler, httpHandler *handler.HttpHandler) {
	r := gin.Default()
	r.Use(middlewares.Authenticate())

	r.GET("/ws", webSocketHandler.HandleWS)

	r.GET("/getConversations/:nextPage", httpHandler.HandlerGetConversations)

	r.Run(":" + os.Getenv("PORT"))
}

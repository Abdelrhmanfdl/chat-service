package router

import (
	handler "chat-chat-go/internal/handlers"
	middlewares "chat-chat-go/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func InitRouter(webSocketHandler handler.WebSocketHandler) {
	r := gin.Default()
	r.Use(middlewares.Authenticate())

	r.GET("/ws", webSocketHandler.HandleWS)

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.Done()
	})

	r.Run(":8082")
}

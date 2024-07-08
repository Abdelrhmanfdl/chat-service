package handlers

import (
	"chat-chat-go/internal/models"
	"chat-chat-go/internal/services"
	"chat-chat-go/internal/websocketmanager"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type WebSocketHandler struct {
	upgrader         websocket.Upgrader
	webSocketManager *websocketmanager.WebSocketManager
	chatService      *services.ChatService
}

func NewWebSocketHandler(chatService *services.ChatService, webSocketManager *websocketmanager.WebSocketManager) *WebSocketHandler {
	return &WebSocketHandler{
		// TODO: Update check origin
		upgrader:         websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }},
		webSocketManager: webSocketManager,
		chatService:      chatService,
	}
}

func (wsh *WebSocketHandler) HandleWS(ctx *gin.Context) {
	conn, err := wsh.upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	userId := ctx.GetString("userId")
	wsh.webSocketManager.AddConnection(userId, conn)
	wsh.chatService.HandleUserConnection(userId)

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received from user %s: %s", userId, message)

		var dtoMessage models.DtoMessage
		if err := json.Unmarshal(message, &dtoMessage); err == nil {
			wsh.chatService.HandleSendMessage(userId, dtoMessage)
			conn.WriteJSON(models.SocketResponse{Succeed: true})
		} else {
			log.Println("Failed to parse message:", err)
			conn.WriteJSON(models.SocketResponse{Succeed: false, Message: "Failed to parse message"})
		}
	}

	wsh.chatService.HandleUserDisconnection(userId)
	wsh.webSocketManager.RemoveConnection(userId)
}

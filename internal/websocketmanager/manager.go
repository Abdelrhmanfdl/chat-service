package websocketmanager

import (
	"chat-chat-go/internal/models"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type SafeWebSocket struct {
	Conn *websocket.Conn
	mu   sync.Mutex
}

type WebSocketManager struct {
	connections map[string]*SafeWebSocket
	// connectionsMutex sync.Mutex
}

func NewWebSocketManager() *WebSocketManager {
	return &WebSocketManager{
		connections: make(map[string]*SafeWebSocket),
	}
}

func (m *WebSocketManager) AddConnection(userId string, conn *websocket.Conn) {
	// m.connectionsMutex.Lock()
	// defer m.connectionsMutex.Unlock()
	m.connections[userId] = &SafeWebSocket{Conn: conn, mu: sync.Mutex{}}
	log.Printf("Connection added for user %s\n", userId)
}

func (m *WebSocketManager) RemoveConnection(userId string) {
	// m.connectionsMutex.Lock()
	// defer m.connectionsMutex.Unlock()
	delete(m.connections, userId)
}

func (m *WebSocketManager) GetSafeWebSocket(userId string) *SafeWebSocket {
	return m.connections[userId]
}

func (m *WebSocketManager) SendMessage(msg models.DtoMessage) {
	if safeSocketConn := m.GetSafeWebSocket(msg.ToUser); safeSocketConn != nil {
		safeSocketConn.mu.Lock()
		defer safeSocketConn.mu.Unlock()
		safeSocketConn.Conn.WriteJSON(msg)
	} else {
		log.Println("User not connected")
	}
}

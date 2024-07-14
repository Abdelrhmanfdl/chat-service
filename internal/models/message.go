package models

import "time"

type Message struct {
	ConversationId string    `json:"conversation_id"`
	MessageId      string    `json:"message_id"`
	Content        string    `json:"content"`
	SenderId       string    `json:"sender_id"`
	ReceiverId     string    `json:"receiver_id"`
	Created_at     time.Time `json:"created_at"`
}

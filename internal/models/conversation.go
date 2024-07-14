package models

import "time"

type Conversation struct {
	SenderId             string    `json:"sender_id"`
	ConversationId       string    `json:"conversation_id"`
	ParticipantId        string    `json:"participant_id"`
	LastMessageId        string    `json:"last_message_id"`
	LastMessageContent   string    `json:"last_message_content"`
	LastMessageTimestamp string    `json:"last_message_timestamp"`
	Created_at           time.Time `json:"created_at"`
}

type Message struct {
	ConversationId string    `json:"conversation_id"`
	MessageId      string    `json:"message_id"`
	Content        string    `json:"content"`
	SenderId       string    `json:"sender_id"`
	ReceiverId     string    `json:"receiver_id"`
	Created_at     time.Time `json:"created_at"`
}

package models

import "time"

type Conversation struct {
	SenderId             string    `json:"sender_id"`
	ConversationId       string    `json:"conversation_id"`
	ParticipantId        string    `json:"participant_id"`
	LastMessageId        string    `json:"last_message_id"`
	LastMessageContent   string    `json:"last_message_content"`
	LastMessageTimestamp time.Time `json:"last_message_timestamp"`
	CreatedAt            time.Time `json:"created_at"`
}

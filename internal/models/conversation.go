package models

import "time"

type Conversation struct {
	Participant1Id       string    `json:"participant1_id"`
	Participant2Id       string    `json:"participant2_id"`
	ConversationId       string    `json:"conversation_id"`
	LastMessageId        string    `json:"last_message_id"`
	LastMessageContent   string    `json:"last_message_content"`
	LastMessageTimestamp time.Time `json:"last_message_timestamp"`
	CreatedAt            time.Time `json:"created_at"`
}

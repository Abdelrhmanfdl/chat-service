package models

import "time"

type ConversationDto struct {
	Participant1         UserDto   `json:"participant1"`
	Participant2         UserDto   `json:"participant2"`
	ConversationId       string    `json:"conversation_id"`
	LastMessageId        string    `json:"last_message_id"`
	LastMessageContent   string    `json:"last_message_content"`
	LastMessageTimestamp time.Time `json:"last_message_timestamp"`
}

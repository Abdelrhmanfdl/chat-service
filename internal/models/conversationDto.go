package models

import "time"

type ConversationDto struct {
	Sender               UserDto   `json:"sender"`
	Participant          UserDto   `json:"participant"`
	ConversationId       string    `json:"conversation_id"`
	LastMessageId        string    `json:"last_message_id"`
	LastMessageContent   string    `json:"last_message_content"`
	LastMessageTimestamp time.Time `json:"last_message_timestamp"`
}

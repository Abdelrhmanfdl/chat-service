package responses

import "chat-chat-go/internal/models"

type Conversations struct {
	Conversations []models.ConversationDto `json:"conversations"`
	NextPage      string                   `json:"nextPage"`
}

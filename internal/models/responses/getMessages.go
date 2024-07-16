package responses

import "chat-chat-go/internal/models"

type Messages struct {
	Messages []models.Message `json:"messages"`
	NextPage string           `json:"nextPage"`
}

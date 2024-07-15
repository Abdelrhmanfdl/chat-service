package chatrepo

import "chat-chat-go/internal/models"

type ChatRepository interface {
	InsertConversation(conversation models.Conversation) (conversationId string, err error)
	GetConversationsBySender(senderId string, paginationState []byte) (conversations []models.Conversation, pageState []byte, err error)

	InsertMessage(message models.Message) (err error)
	GetMessagesByConversation(conversationId string, lastPageState []byte) (messages []models.Message, pageState []byte, err error)
	GetMessageById(messageId string) (messages []models.Message, err error)
}

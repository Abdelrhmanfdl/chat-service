package chatrepo

import (
	"chat-chat-go/internal/models"
	"time"

	"github.com/gocql/gocql"
)

type scyllaConversation struct {
	Participant1Id       gocql.UUID `json:"participant1_id"`
	Participant2Id       gocql.UUID `json:"participant2_id"`
	ConversationId       gocql.UUID `json:"conversation_id"`
	LastMessageId        gocql.UUID `json:"last_message_id"`
	LastMessageContent   string     `json:"last_message_content"`
	LastMessageTimestamp time.Time  `json:"last_message_timestamp"`
	CreatedAt            time.Time  `json:"created_at"`
}

func createScyllaConversation(conversation models.Conversation) (newScyllaConversation *scyllaConversation, err error) {
	newScyllaConversation = &scyllaConversation{}
	newScyllaConversation.Participant1Id, err = gocql.ParseUUID(conversation.Participant1Id)
	if err != nil {
		return nil, err
	}
	newScyllaConversation.Participant2Id, err = gocql.ParseUUID(conversation.Participant2Id)

	if err != nil {
		return nil, err
	}

	newScyllaConversation.ConversationId = gocql.TimeUUID()
	newScyllaConversation.CreatedAt = time.Now()

	return newScyllaConversation, nil
}

func (c *scyllaConversation) ConvertToVanillaConversation() *models.Conversation {
	return &models.Conversation{
		Participant1Id:       c.Participant1Id.String(),
		Participant2Id:       c.Participant2Id.String(),
		ConversationId:       c.ConversationId.String(),
		LastMessageId:        c.LastMessageId.String(),
		LastMessageContent:   c.LastMessageContent,
		LastMessageTimestamp: c.LastMessageTimestamp,
		CreatedAt:            c.CreatedAt,
	}
}

type scyllaMessage struct {
	ConversationId gocql.UUID `json:"conversation_id"`
	MessageId      gocql.UUID `json:"message_id"`
	SenderId       gocql.UUID `json:"sender_id"`
	ReceiverId     gocql.UUID `json:"receiver_id"`
	Content        string     `json:"content"`
	CreatedAt      time.Time  `json:"created_at"`
}

func createScyllaMessage(message models.Message) (newScyllaMessage *scyllaMessage, err error) {
	newScyllaMessage = &scyllaMessage{}
	newScyllaMessage.SenderId, err = gocql.ParseUUID(message.SenderId)
	if err != nil {
		return nil, err
	}

	newScyllaMessage.ReceiverId, err = gocql.ParseUUID(message.ReceiverId)
	if err != nil {
		return nil, err
	}

	newScyllaMessage.ConversationId, err = gocql.ParseUUID(message.ConversationId)
	if err != nil {
		return nil, err
	}

	newScyllaMessage.Content = message.Content
	newScyllaMessage.MessageId = gocql.TimeUUID()
	newScyllaMessage.CreatedAt = time.Now()

	return newScyllaMessage, nil
}

func (m *scyllaMessage) ConvertToVanillaMessage() *models.Message {
	return &models.Message{
		SenderId:       m.SenderId.String(),
		ReceiverId:     m.ReceiverId.String(),
		ConversationId: m.ConversationId.String(),
		MessageId:      m.MessageId.String(),
		Content:        m.Content,
		CreatedAt:      m.CreatedAt,
	}
}

package chatrepo

import (
	"chat-chat-go/internal/models"
	"log"

	"github.com/gocql/gocql"
)

type ScyllaChatRepository struct {
	session  *gocql.Session
	pageSize int
}

func NewScyllaChatRepository(scyllaURL string) *ScyllaChatRepository {
	cluster := gocql.NewCluster(scyllaURL)
	cluster.Keyspace = "chatchatgo"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Unable to connect to ScyllaDB:", err)
	}
	return &ScyllaChatRepository{session: session, pageSize: 3}
}

func (r *ScyllaChatRepository) InsertConversation(conversation models.Conversation) (conversationId string, err error) {
	newConv, err := createScyllaConversation(conversation)
	if err != nil {
		log.Println("can not create new conversation:", err)
	}
	query := `INSERT INTO conversations_by_user (sender_id, conversation_id, participant_id, created_at) VALUES(?, ?, ?, ?)`
	err = r.session.Query(query, newConv.SenderId, newConv.ConversationId, newConv.ParticipantId, newConv.CreatedAt, query).Exec()
	return newConv.ConversationId.String(), err
}

func (r *ScyllaChatRepository) GetConversationsBySender(senderId string, lastPageState []byte) (conversations []models.Conversation, pageState []byte, err error) {
	queryStr := `SELECT sender_id, conversation_id, participant_id, last_message_id, last_message_content, last_message_timestamp 
				 FROM conversations_by_user
				 WHERE sender_id = ?`

	queryItr := r.session.Query(queryStr, senderId).PageSize(r.pageSize).PageState(lastPageState).Iter()

	conversation := models.Conversation{}

	for queryItr.Scan(&conversation.SenderId, &conversation.ConversationId, &conversation.ParticipantId,
		&conversation.LastMessageContent, &conversation.LastMessageTimestamp) {
		conversations = append(conversations, conversation)
	}

	if err = queryItr.Close(); err != nil {
		return nil, nil, err
	}

	return conversations, queryItr.PageState(), err
}

func (r *ScyllaChatRepository) InsertMessage(message models.Message) (err error) {
	newMessage, err := createScyllaMessage(message)
	if err != nil {
		log.Println("can not create new conversation:", err)
	}

	batch := r.session.NewBatch(gocql.LoggedBatch)
	queryToMessages := `INSERT INTO messages (sender_id, receiver_id, conversation_id, message_id, content, created_at) VALUES(?, ?, ?, ?)`
	batch.Query(queryToMessages, newMessage.SenderId, newMessage.ReceiverId, newMessage.ConversationId, newMessage.MessageId, newMessage.Content, newMessage.CreatedAt)
	queryToMessagesByConversation := `INSERT INTO messages_by_conversation (sender_id, receiver_id, conversation_id, message_id, content, created_at) VALUES(?, ?, ?, ?)`
	batch.Query(queryToMessagesByConversation, newMessage.SenderId, newMessage.ReceiverId, newMessage.ConversationId, newMessage.MessageId, newMessage.Content, newMessage.CreatedAt)

	return r.session.ExecuteBatch(batch)
}

func (r *ScyllaChatRepository) GetMessagesByConversation(conversationId string, lastPageState []byte) (messages []models.Message, pageState []byte, err error) {
	queryStr := `SELECT sender_id, receiver_id, conversation_id, message_id, content, created_at 
				 FROM messages_by_conversation
				 WHERE conversation_id = ?`

	queryItr := r.session.Query(queryStr, conversationId).PageSize(r.pageSize).PageState(lastPageState).Iter()

	message := models.Message{}

	for queryItr.Scan(&message.SenderId, &message.ReceiverId, &message.ConversationId,
		&message.MessageId, &message.Content, &message.CreatedAt) {
		messages = append(messages, message)
	}

	if err = queryItr.Close(); err != nil {
		return nil, nil, err
	}

	return messages, queryItr.PageState(), err
}

func (r *ScyllaChatRepository) GetMessageById(messageId string) (messages []models.Message, err error) {
	queryStr := `SELECT sender_id, receiver_id, conversation_id, message_id, content, created_at 
				 FROM messages
				 WHERE message_id = ?`

	message := models.Message{}

	err = r.session.Query(queryStr, messageId).Scan(&message.SenderId, &message.ReceiverId, &message.ConversationId,
		&message.MessageId, &message.Content, &message.CreatedAt)

	if err != nil {
		return nil, err
	}

	return messages, err
}

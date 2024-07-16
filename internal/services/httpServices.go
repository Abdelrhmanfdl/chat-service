package services

import (
	"chat-chat-go/internal/errs"
	"chat-chat-go/internal/models"
)

func (chatService *ChatService) fillConversationDtosByUsersData(conversations []models.Conversation) (conversationDtos []models.ConversationDto, err error) {
	var othersIds []string = make([]string, len(conversationDtos))
	conversationMap := make(map[string]int)

	for idx, conversation := range conversations {
		othersIds = append(othersIds, conversation.Participant1Id)
		conversationMap[conversation.Participant1Id] = idx
	}

	othersDtos, err := chatService.userService.GetUsersData(othersIds)
	if err != nil {
		return nil, err
	}

	for idx, otherDto := range othersDtos {
		conversation := conversations[conversationMap[otherDto.ID]]
		conversationDtos = append(conversationDtos, models.ConversationDto{
			Participant1:         models.UserDto{ID: conversation.Participant1Id}, // complete
			Participant2:         models.UserDto{ID: otherDto.ID, Username: otherDto.Username},
			ConversationId:       conversation.ConversationId,
			LastMessageId:        conversation.LastMessageId,
			LastMessageContent:   conversation.LastMessageContent,
			LastMessageTimestamp: conversation.LastMessageTimestamp,
		})

		othersIds = append(othersIds, conversation.Participant1Id)
		conversationMap[conversation.Participant1Id] = idx
	}

	return conversationDtos, err
}

func (chatService *ChatService) GetConversations(userId string, nextPage string) (conversationDtos []models.ConversationDto, newNextPage string, err error) {
	conversations, newNextPageBytes, err := chatService.chatRepository.GetConversationsByUser(userId, []byte(nextPage))
	if err != nil {
		return nil, "", err
	}

	conversationDtos, err = chatService.fillConversationDtosByUsersData(conversations)
	if err != nil {
		return nil, "", err
	}

	newNextPage = string(newNextPageBytes)
	return conversationDtos, newNextPage, nil
}

func (chatService *ChatService) GetMessagesByConversation(userId string, conversationId string, nextPage string) (messages []models.Message, newNextPage string, err error) {
	messages, newNextPageBytes, err := chatService.chatRepository.GetMessagesByConversation(conversationId, []byte(nextPage))
	if err != nil {
		return nil, "", err
	}

	if len(messages) > 0 && (messages[0].ReceiverId == userId || messages[0].SenderId == userId) {
		return nil, "", &errs.UnauthorizedFetchError{Message: "unauthorized fetching messages"}
	}

	return messages, string(newNextPageBytes), nil
}

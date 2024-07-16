package handlers

import (
	"chat-chat-go/internal/models/responses"
	"chat-chat-go/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	chatService *services.ChatService
}

func NewHttpHandler(chatService *services.ChatService) *HttpHandler {
	return &HttpHandler{
		chatService: chatService,
	}
}

func (h *HttpHandler) HandlerGetConversations(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	nextPage := ctx.Param("nextPage")

	conversations, nextPage, err := h.chatService.GetConversations(userId, nextPage)

	if err != nil {
		// TODO: classify error
		ctx.Status(http.StatusInternalServerError)
	}

	ctx.JSON(http.StatusOK, responses.Conversations{Conversations: conversations, NextPage: nextPage})
}

package handlers

import (
	"chat-chat-go/internal/errs"
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

func (h *HttpHandler) HandlerGetMessagesByConversation(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	nextPage := ctx.Request.URL.Query().Get("nextPage")
	conversationId := ctx.Request.URL.Query().Get("conversationId")

	if conversationId == "" {
		// TODO: classify error
		ctx.JSON(http.StatusBadRequest, &errs.MissingQueryParamsErros{Message: "conversationId must be defined"})
	}

	messages, nextPage, err := h.chatService.GetMessagesByConversation(userId, conversationId, nextPage)
	if err != nil {
		if _, ok := err.(*errs.UnauthorizedFetchError); !ok {
			ctx.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
			return
		} else {
			// TODO: classify error
			ctx.Status(http.StatusInternalServerError)
		}
	}

	ctx.JSON(http.StatusOK, responses.Messages{Messages: messages, NextPage: nextPage})
}

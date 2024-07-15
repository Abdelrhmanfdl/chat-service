package userrepo

import "chat-chat-go/internal/models"

type UserRepository interface {
	GetUserById(id string) (user *models.User, err error)
}

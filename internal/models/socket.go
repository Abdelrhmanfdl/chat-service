package models

type DtoInSocketMessage interface{}
type DtoOutSocketMessage interface{}

type DtoResponseSocketResponse struct {
	Succeed bool
	Message string
}

type DtoInChatSocketMessage struct {
	FromUser string
	ToUser   string
	Content  string
}

type DtoOutChatSocketMessage struct {
	FromUser string
	ToUser   string
	Content  string
}

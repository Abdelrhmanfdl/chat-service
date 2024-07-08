package models

type SocketResponse struct {
	Succeed bool
	Message string
}

type DtoMessage struct {
	FromUser string
	ToUser   string
	Content  string
}

package userregistry

type UserRegistry interface {
	GetUserRegistry(userId string) (string, error)
	RegisterUser(userId, registry string) error
	UnregisterUser(userId string) error
	Close() error
}

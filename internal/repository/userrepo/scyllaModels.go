package userrepo

import "github.com/gocql/gocql"

type scyllaUser struct {
	ID        gocql.UUID `json:"id"`
	Username  string     `json:"name"`
	Email     string     `json:"email"`
	Password  string     `json:"password"`
	CreatedAt int        `json:"createdAt"`
}

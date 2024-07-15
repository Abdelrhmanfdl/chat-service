package userrepo

import (
	"chat-chat-go/internal/models"
	"log"

	"github.com/gocql/gocql"
)

type ScyllaUserRepository struct {
	session *gocql.Session
}

func NewScyllaUserRepository(scyllaURL string) *ScyllaUserRepository {
	cluster := gocql.NewCluster(scyllaURL)
	cluster.Keyspace = "chatchatgo"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal("Unable to connect to ScyllaDB:", err)
	}
	return &ScyllaUserRepository{session: session}
}

func (s *ScyllaUserRepository) GetUserById(id string) (user *models.User, err error) {
	var userScylla scyllaUser
	query := `SELECT id, username, email FROM users WHERE id = ?`
	if err := s.session.Query(query, id).Scan(&userScylla.ID, &userScylla.Username, &userScylla.Email); err != nil {
		return nil, classifyError(err)
	}
	return &models.User{
		ID:       userScylla.ID.String(),
		Username: userScylla.Username,
		Email:    userScylla.Email,
	}, nil
}

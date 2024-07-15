package userrepo

import (
	"chat-chat-go/internal/errs"

	"github.com/gocql/gocql"
)

func classifyError(err error) error {
	if err == gocql.ErrNotFound {
		return &errs.NotFoundError{Message: "not found"}
	}
	if err, ok := err.(gocql.RequestError); ok && err.Code() == gocql.ErrCodeUnavailable {
		return &errs.ConnectionError{Message: err.Error()}
	}
	if err, ok := err.(gocql.RequestErrReadTimeout); ok && err.Code() == gocql.ErrCodeReadTimeout {
		return &errs.TimeoutError{Message: err.Error()}
	}
	return err
}

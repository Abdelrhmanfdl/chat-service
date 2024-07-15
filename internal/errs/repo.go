package errs

import "fmt"

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return "Not found"
}

type ConnectionError struct {
	Message string
}

func (e *ConnectionError) Error() string {
	return fmt.Sprintf("connection error: %s", e.Message)
}

type TimeoutError struct {
	Message string
}

func (e *TimeoutError) Error() string {
	return fmt.Sprintf("timeout error: %s", e.Message)
}

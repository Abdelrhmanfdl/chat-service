package errs

import "fmt"

type MissingQueryParamsErros struct {
	Message string
}

func (e *MissingQueryParamsErros) Error() string {
	return fmt.Sprintf("missing query params: %s", e.Message)
}

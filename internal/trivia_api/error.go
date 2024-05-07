package triviaapi

import "fmt"

type APIError struct {
	Code int
}

func (e *APIError) Error() string {
	return fmt.Sprintf("opentdb:API error with code:%d ,check docs for more info", e.Code)
}
func NewAPIError(code int) *APIError {
	return &APIError{
		Code: code,
	}
}

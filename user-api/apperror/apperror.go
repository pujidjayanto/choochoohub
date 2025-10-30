package apperror

import (
	"fmt"
)

type AppError struct {
	StatusCode int
	ErrorCode  int
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewAppError(status, errCode int, err error) *AppError {
	return &AppError{
		StatusCode: status,
		ErrorCode:  errCode,
		Err:        err,
		Message:    err.Error(),
	}
}

func AsAppError(err error) (*AppError, bool) {
	appErr, ok := err.(*AppError)
	return appErr, ok
}

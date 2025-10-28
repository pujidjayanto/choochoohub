package apperror

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type AppError struct {
	StatusCode int
	Message    string
	Err        error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Err)
	}
	return e.Message
}

func NewAppError(status int, msg string, err error) *AppError {
	return &AppError{
		StatusCode: status,
		Message:    msg,
		Err:        err,
	}
}

func AsAppError(err error) (*AppError, bool) {
	appErr, ok := err.(*AppError)
	return appErr, ok
}

func NewBadRequestError(err error) *AppError {
	return &AppError{
		StatusCode: fiber.StatusBadRequest,
		Message:    "bad request",
		Err:        err,
	}
}

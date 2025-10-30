package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/apperror"
)

func BindAndValidate(c echo.Context, i any) error {
	if err := c.Bind(i); err != nil {
		return apperror.NewAppError(http.StatusBadRequest, apperror.CodeBadRequest, err)
	}

	if err := c.Validate(i); err != nil {
		return apperror.NewAppError(http.StatusBadRequest, apperror.CodeBadRequest, err)
	}

	return nil
}

package delivery

import (
	"net/http"
	"runtime/debug"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/apperror"
)

func SuccessNoContent(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func SuccessCreated(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}

func Success(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, SuccessResponse{
		Message: SuccessMessage,
		Data:    data,
	})
}

func Failed(c echo.Context, err error) error {
	var appErr *apperror.AppError
	if e, ok := apperror.AsAppError(err); ok {
		appErr = e
	} else {
		appErr = apperror.NewAppError(
			http.StatusInternalServerError,
			apperror.CodeInternalServerError,
			err,
		)
	}

	resp := ErrorResponse{
		Error:     appErr.Message,
		ErrorCode: appErr.ErrorCode,
	}

	// Optional trace
	if header := c.Request().Header.Get("X-Enable-Trace"); header != "" {
		if isAppTrace, parseErr := strconv.ParseBool(header); parseErr == nil && isAppTrace {
			resp.Trace = string(debug.Stack())
		}
	}

	c.Set("Connection", "close")

	return c.JSON(appErr.StatusCode, resp)
}

package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

type SignInController interface {
	SignIn(c echo.Context) error
}

type signInController struct {
	usecase usecase.SignInUsecase
}

func NewSignInController(usecase usecase.SignInUsecase) SignInController {
	return &signInController{usecase}
}

func (signInController *signInController) SignIn(c echo.Context) error {
	var req dto.SigninRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	resp, err := signInController.usecase.SignIn(c.Request().Context(), req)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

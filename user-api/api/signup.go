package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

type SignUpController interface {
	SignUp(c echo.Context) error
}

type signUpController struct {
	usecase usecase.SignUpUsecase
}

func NewSignUpController(usecase usecase.SignUpUsecase) SignUpController {
	return &signUpController{usecase}
}

func (signUpController *signUpController) SignUp(c echo.Context) error {
	var req dto.SignupRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err := signUpController.usecase.Create(c.Request().Context(), req)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

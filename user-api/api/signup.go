package api

import (
	"github.com/labstack/echo/v4"
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
	return nil
}

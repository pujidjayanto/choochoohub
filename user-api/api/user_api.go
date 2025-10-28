package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

type UserApi interface {
	SignIn(c echo.Context) error
	SignUp(c echo.Context) error
}

type userApi struct {
	signInUsecase usecase.SignInUsecase
	signUpUsecase usecase.SignUpUsecase
}

func NewUserApi(signInUsecase usecase.SignInUsecase, signUpUsecase usecase.SignUpUsecase) UserApi {
	return &userApi{signInUsecase, signUpUsecase}
}

func (ua *userApi) SignIn(c echo.Context) error {
	var req dto.SigninRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	resp, err := ua.signInUsecase.SignIn(c.Request().Context(), req)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.JSON(http.StatusOK, resp)
}

func (ua *userApi) SignUp(c echo.Context) error {
	var req dto.SignupRequest
	if err := c.Bind(&req); err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err := ua.signUpUsecase.Create(c.Request().Context(), req)
	if err != nil {
		return c.String(http.StatusUnprocessableEntity, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

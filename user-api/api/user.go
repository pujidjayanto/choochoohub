package api

import (
	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/delivery"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

type UserApi interface {
	SignUp(c echo.Context) error
}

type userApi struct {
	userUsecase usecase.UserUsecase
}

func NewUserApi(userUsecase usecase.UserUsecase) UserApi {
	return &userApi{userUsecase}
}

func (ua *userApi) SignUp(c echo.Context) error {
	var req dto.SignupRequest
	if err := BindAndValidate(c, &req); err != nil {
		return delivery.Failed(c, err)
	}

	err := ua.userUsecase.Signup(c.Request().Context(), req)
	if err != nil {
		return delivery.Failed(c, err)
	}

	return delivery.SuccessNoContent(c)
}

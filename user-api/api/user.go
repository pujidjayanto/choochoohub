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
	signUpUsecase usecase.SignUpUsecase
}

func NewUserApi(signUpUsecase usecase.SignUpUsecase) UserApi {
	return &userApi{signUpUsecase}
}

func (ua *userApi) SignUp(c echo.Context) error {
	var req dto.SignupRequest
	if err := BindAndValidate(c, &req); err != nil {
		return delivery.Failed(c, err)
	}

	err := ua.signUpUsecase.Create(c.Request().Context(), req)
	if err != nil {
		return delivery.Failed(c, err)
	}

	return delivery.SuccessNoContent(c)
}

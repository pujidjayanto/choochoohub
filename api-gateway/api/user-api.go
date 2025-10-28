package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pujidjayanto/choochoohub/api-gateway/apperror"
	userapi "github.com/pujidjayanto/choochoohub/api-gateway/client/user-api"
	"github.com/pujidjayanto/choochoohub/api-gateway/dto"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/delivery"
)

type UserApi interface {
	Signup(c *fiber.Ctx) error
	Signin(c *fiber.Ctx) error
}

type userApi struct {
	client userapi.Client
}

func NewUserApi(client userapi.Client) UserApi {
	return &userApi{client: client}
}

func (userApi *userApi) Signin(c *fiber.Ctx) error {
	var req dto.SigninRequest
	if err := c.BodyParser(&req); err != nil {
		appErr := apperror.NewBadRequestError(err)
		return delivery.Failed(c, appErr.StatusCode, appErr.Message)
	}

	resp, err := userApi.client.Signin(c.UserContext(), &userapi.SigninRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return delivery.Failed(c, err.StatusCode, err.Error())
	}

	return delivery.Success(c, resp)
}

func (userApi *userApi) Signup(c *fiber.Ctx) error {
	var req dto.SignupRequest
	if err := c.BodyParser(&req); err != nil {
		appErr := apperror.NewBadRequestError(err)
		return delivery.Failed(c, appErr.StatusCode, appErr.Message)
	}

	err := userApi.client.Signup(c.UserContext(), &userapi.SignupRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return delivery.Failed(c, err.StatusCode, err.Error())
	}

	return nil
}

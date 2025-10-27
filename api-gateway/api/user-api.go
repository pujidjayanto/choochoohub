package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	userapi "github.com/pujidjayanto/choochoohub/api-gateway/client/user-api"
	"github.com/pujidjayanto/choochoohub/api-gateway/dto"
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
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("bad request")
	}

	err := userApi.client.Signin(c.UserContext(), &userapi.SigninRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return nil
}

func (userApi *userApi) Signup(c *fiber.Ctx) error {
	var req dto.SignupRequest
	if err := c.BodyParser(req); err != nil {
		return c.Status(http.StatusBadRequest).SendString("bad request")
	}

	err := userApi.client.Signup(c.UserContext(), &userapi.SignupRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return nil
}

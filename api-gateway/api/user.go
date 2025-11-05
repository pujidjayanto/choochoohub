package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/pujidjayanto/choochoohub/api-gateway/appctx"
	"github.com/pujidjayanto/choochoohub/api-gateway/apperror"
	userapi "github.com/pujidjayanto/choochoohub/api-gateway/client/user-api"
	"github.com/pujidjayanto/choochoohub/api-gateway/dto"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/delivery"
)

type UserApi interface {
	Signup(c *fiber.Ctx) error
}

type userApi struct {
	client userapi.Client
}

func NewUserApi(client userapi.Client) UserApi {
	return &userApi{client: client}
}

func (userApi *userApi) Signup(c *fiber.Ctx) error {
	var req dto.SignupRequest
	if err := c.BodyParser(&req); err != nil {
		appErr := apperror.NewAppError(http.StatusBadRequest, apperror.CodeInternalBadRequest, err)
		return delivery.Failed(c, appErr.StatusCode, appErr.Error())
	}

	reqID := c.Locals("requestid").(string)
	ctx := appctx.WithRequestID(c.UserContext(), reqID)
	err := userApi.client.Signup(ctx, &userapi.SignupRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		return delivery.Failed(c, err.StatusCode, err.Error())
	}

	return nil
}

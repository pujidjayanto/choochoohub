package userapi

import (
	"context"
	"net/http"

	"github.com/pujidjayanto/choochoohub/api-gateway/appctx"
	"github.com/pujidjayanto/choochoohub/api-gateway/apperror"
	"github.com/pujidjayanto/choochoohub/api-gateway/client/config"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/httpclient"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/urlbuilder"
)

type Client interface {
	Signin(ctx context.Context, req *SigninRequest) (*SigninResponse, *apperror.AppError)
	Signup(ctx context.Context, req *SignupRequest) *apperror.AppError
}

type client struct {
	httpClient httpclient.Client
	host       string
	port       string
}

func NewClient(httpclient httpclient.Client, config config.UserApi) Client {
	return &client{httpClient: httpclient, host: config.Host, port: config.Port}
}

func (client *client) Signin(ctx context.Context, req *SigninRequest) (*SigninResponse, *apperror.AppError) {
	var resp SigninResponse

	url, err := urlbuilder.Build(client.host, client.port, "/v1/signin")
	if err != nil {
		return nil, apperror.NewAppError(http.StatusUnprocessableEntity, "user-api signin build url failed", err)
	}

	requestID := appctx.RequestID(ctx)

	headers := map[string]string{}
	if requestID != "" {
		headers["X-Request-ID"] = requestID
	}

	err = client.httpClient.FireRequest(ctx, http.MethodPost, url, headers, req, &resp)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusUnprocessableEntity, "user-api signin request failed", err)
	}

	return &resp, nil
}

func (client *client) Signup(ctx context.Context, req *SignupRequest) *apperror.AppError {
	url, err := urlbuilder.Build(client.host, client.port, "/v1/signup")
	if err != nil {
		return apperror.NewAppError(http.StatusUnprocessableEntity, "user-api signup build url failed", err)
	}

	requestID := appctx.RequestID(ctx)

	headers := map[string]string{}
	if requestID != "" {
		headers["X-Request-ID"] = requestID
	}

	err = client.httpClient.FireRequest(ctx, http.MethodPost, url, headers, req, nil)
	if err != nil {
		return apperror.NewAppError(http.StatusUnprocessableEntity, "user-api signup request failed", err)
	}

	return nil
}

package userapi

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/pujidjayanto/choochoohub/api-gateway/appctx"
	"github.com/pujidjayanto/choochoohub/api-gateway/apperror"
	"github.com/pujidjayanto/choochoohub/api-gateway/client/config"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/httpclient"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/jsonb"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/urlbuilder"
)

type ApiErrorResponse struct {
	Error     string `json:"error"`
	ErrorCode string `json:"errorCode"`
	Trace     string `json:"trace,omitempty"`
}

type Client interface {
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

func (client *client) Signup(ctx context.Context, req *SignupRequest) *apperror.AppError {
	url, err := urlbuilder.Build(client.host, client.port, SignUpPath)
	if err != nil {
		return apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, errors.New("failed to build url"))
	}

	requestID := appctx.RequestID(ctx)

	headers := map[string]string{}
	if requestID != "" {
		headers["X-Request-ID"] = requestID
	}

	resp, err := client.httpClient.Fire(ctx, http.MethodPost, url, headers, req)
	if err != nil {
		return apperror.NewAppError(
			http.StatusServiceUnavailable,
			apperror.CodeUserApiInternalServerError,
			err,
		)
	}

	if resp.StatusCode >= 400 {
		var apiError ApiErrorResponse
		if err := jsonb.Unmarshal(resp.Body, &apiError); err != nil {
			return apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
		}

		errCode, _ := strconv.Atoi(apiError.ErrorCode)
		return apperror.NewAppError(http.StatusUnprocessableEntity, errCode, errors.New(apiError.Error))
	}

	// expect no content
	return nil
}

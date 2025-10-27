package userapi

import (
	"context"
	"net/http"

	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/httpclient"
)

type Client interface {
	Signin(ctx context.Context, req *SigninRequest) error
	Signup(ctx context.Context, req *SignupRequest) error
}

type client struct {
	httpClient httpclient.Client
	host       string
}

func NewClient(httpclient httpclient.Client) Client {
	return &client{httpClient: httpclient}
}

func (client *client) Signin(ctx context.Context, req *SigninRequest) error {
	var resp SigninResponse
	url := client.host + "/v1/signin"
	err := client.httpClient.FireRequest(ctx, http.MethodPost, url, nil, nil, req, &resp)
	if err != nil {
		return err
	}

	return nil
}

func (client *client) Signup(ctx context.Context, req *SignupRequest) error {
	url := client.host + "/v1/signup"
	err := client.httpClient.FireRequest(ctx, http.MethodPost, url, nil, nil, req, nil)
	if err != nil {
		return err
	}

	return nil
}

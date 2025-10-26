package userapi

import (
	"context"

	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/httpclient"
)

type Client interface {
	Signin(ctx context.Context, req *SigninRequest) error
	Signup(ctx context.Context, req *SignupRequest) error
}

type client struct {
	httpClient httpclient.Client
}

func NewClient(httpclient httpclient.Client) Client {
	return &client{httpClient: httpclient}
}

func (client *client) Signin(ctx context.Context, req *SigninRequest) error {
	return nil
}

func (client *client) Signup(ctx context.Context, req *SignupRequest) error {
	return nil
}

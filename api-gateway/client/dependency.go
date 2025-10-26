package client

import (
	userapi "github.com/pujidjayanto/choochoohub/api-gateway/client/user-api"
	"github.com/pujidjayanto/choochoohub/api-gateway/pkg/httpclient"
	"github.com/sirupsen/logrus"
)

type Dependency struct {
	UserApiClient userapi.Client
}

func NewDependency(httpclient httpclient.Client, log *logrus.Logger) Dependency {
	return Dependency{
		UserApiClient: userapi.NewClient(httpclient),
	}
}

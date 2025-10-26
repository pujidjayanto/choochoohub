package api

import (
	"github.com/pujidjayanto/choochoohub/api-gateway/client"
)

type Dependency struct {
	UserApi UserApi
}

func NewDependency(externalClient client.Dependency) Dependency {
	return Dependency{
		UserApi: NewUserApi(&externalClient.UserApiClient),
	}
}

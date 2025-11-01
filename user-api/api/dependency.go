package api

import "github.com/pujidjayanto/choochoohub/user-api/usecase"

type Dependency struct {
	PingApi PingApi
	UserApi UserApi
	OtpApi  OtpApi
}

func NewDependency(usecases usecase.Dependency) Dependency {
	return Dependency{
		PingApi: NewPingApi(),
		UserApi: NewUserApi(usecases.SignUpUsecase),
		OtpApi:  NewOtpApi(usecases.OtpUsecase),
	}
}

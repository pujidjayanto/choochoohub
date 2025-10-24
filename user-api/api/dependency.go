package api

import "github.com/pujidjayanto/choochoohub/user-api/usecase"

type Dependency struct {
	PingController   PingController
	SignUpController SignUpController
}

func NewDependency(usecases usecase.Dependency) Dependency {
	return Dependency{
		PingController:   NewPingController(),
		SignUpController: NewSignUpController(usecases.SignUpUsecase),
	}
}

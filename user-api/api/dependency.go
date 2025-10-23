package api

import "github.com/pujidjayanto/choochoohub/user-api/usecase"

type Dependency struct {
	SignUpController SignUpController
}

func NewDependency(usecases usecase.Dependency) Dependency {
	return Dependency{
		SignUpController: NewSignUpController(usecases.SignUpUsecase),
	}
}

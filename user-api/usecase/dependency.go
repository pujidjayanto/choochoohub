package usecase

import "github.com/pujidjayanto/choochoohub/user-api/repository"

type Dependency struct {
	SignUpUsecase SignUpUsecase
}

func NewDependency(repositories repository.Dependency) Dependency {
	return Dependency{
		SignUpUsecase: NewSignupUsecase(repositories.UserRepository),
	}
}

package usecase

import "github.com/pujidjayanto/choochoohub/user-api/repository"

type Dependency struct {
	SignUpUsecase SignUpUsecase
	SignInUsecase SignInUsecase
}

func NewDependency(repositories repository.Dependency) Dependency {
	return Dependency{
		SignUpUsecase: NewSignupUsecase(repositories.UserRepository),
		SignInUsecase: NewSignInUsecase(repositories.UserRepository),
	}
}

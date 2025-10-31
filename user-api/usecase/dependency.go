package usecase

import (
	"github.com/pujidjayanto/choochoohub/user-api/pkg/eventbus"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
)

type Dependency struct {
	SignUpUsecase SignUpUsecase
	OtpUsecase    OtpUsecase
}

func NewDependency(repositories repository.Dependency, eventBus eventbus.EventBus) Dependency {
	return Dependency{
		SignUpUsecase: NewSignupUsecase(repositories.UserRepository, eventBus),
		OtpUsecase:    NewOtpUsecase(repositories.UserOtpRepository),
	}
}

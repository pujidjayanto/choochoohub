package usecase

import (
	"github.com/pujidjayanto/choochoohub/user-api/pkg/eventbus"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
)

type Dependency struct {
	OtpUsecase  OtpUsecase
	UserUsecase UserUsecase
}

func NewDependency(repositories repository.Dependency, eventBus eventbus.EventBus) Dependency {
	return Dependency{
		OtpUsecase:  NewOtpUsecase(repositories.UserOtpRepository, eventBus),
		UserUsecase: NewUserUsecase(repositories.UserRepository, eventBus),
	}
}

package usecase

import (
	"context"

	"github.com/pujidjayanto/choochoohub/user-api/repository"
)

type SignUpUsecase interface {
	Create(c context.Context) error
}

type signUpUsecase struct {
	userRepository repository.UserRepository
}

func NewSignupUsecase(userRepository repository.UserRepository) SignUpUsecase {
	return &signUpUsecase{
		userRepository: userRepository,
	}
}

func (su *signUpUsecase) Create(c context.Context) error {
	return nil
}

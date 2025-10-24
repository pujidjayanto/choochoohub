package usecase

import (
	"context"

	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/pwd"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
)

type SignUpUsecase interface {
	Create(c context.Context, req dto.SignupRequest) error
}

type signUpUsecase struct {
	userRepository repository.UserRepository
}

func NewSignupUsecase(userRepository repository.UserRepository) SignUpUsecase {
	return &signUpUsecase{
		userRepository: userRepository,
	}
}

func (signUpUsecase *signUpUsecase) Create(c context.Context, req dto.SignupRequest) error {
	hashedPwd, err := pwd.Hash(req.Password)
	if err != nil {
		return err
	}

	err = signUpUsecase.userRepository.Create(c, &model.User{PasswordHash: hashedPwd, Email: req.Email})
	if err != nil {
		return err
	}

	return nil
}

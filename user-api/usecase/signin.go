package usecase

import (
	"context"
	"errors"

	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/pwd"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
)

type SignInUsecase interface {
	SignIn(c context.Context, req dto.SigninRequest) (*dto.SigninResponse, error)
}

type siginUsecase struct {
	userRepository repository.UserRepository
}

func NewSignInUsecase(userRepository repository.UserRepository) SignInUsecase {
	return &siginUsecase{
		userRepository: userRepository,
	}
}

func (siginUsecase *siginUsecase) SignIn(c context.Context, req dto.SigninRequest) (*dto.SigninResponse, error) {
	user, err := siginUsecase.userRepository.FindByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}

	validPwd := pwd.Match(user.PasswordHash, req.Password)
	if !validPwd {
		return nil, errors.New("invalid password")
	}

	return &dto.SigninResponse{
		ID:    user.ID.String(),
		Email: user.Email,
	}, nil
}

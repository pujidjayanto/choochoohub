package usecase

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/pujidjayanto/choochoohub/user-api/apperror"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/eventbus"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/stringhash"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
	"gorm.io/gorm"
)

type UserUsecase interface {
	Signup(c context.Context, req dto.SignupRequest) error
	Signin(c context.Context, req dto.SigninRequest) (*dto.SigninResponse, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
	eventBus       eventbus.EventBus
}

func NewUserUsecase(userRepository repository.UserRepository, eventBus eventbus.EventBus) UserUsecase {
	return &userUsecase{userRepository, eventBus}
}

func (u *userUsecase) Signup(c context.Context, req dto.SignupRequest) error {
	hashedPwd, err := stringhash.Hash(req.Password)
	if err != nil {
		return apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
	}

	user, err := u.userRepository.Create(c, &model.User{PasswordHash: hashedPwd, Email: req.Email})
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return apperror.NewAppError(http.StatusUnprocessableEntity, apperror.CodeValidationFailed, err)
		}
		return apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
	}

	u.eventBus.Publish("user.created", dto.OtpRequest{
		UserId:      user.ID,
		Channel:     string(model.UserOtpChannelEmail),
		Destination: user.Email,
		Purpose:     string(model.UserOtpPurposeSignup),
		ExpiredAt:   time.Now().Add(5 * time.Minute),
	})

	return nil
}

func (u *userUsecase) Signin(c context.Context, req dto.SigninRequest) (*dto.SigninResponse, error) {
	user, err := u.userRepository.FindByEmail(c, req.Email)
	if err != nil {
		return nil, apperror.NewAppError(http.StatusUnauthorized, apperror.CodeUnauthorized, err)
	}

	isMatched := stringhash.Match(user.PasswordHash, req.Password)
	if !isMatched {
		return nil, apperror.NewAppError(http.StatusUnauthorized, apperror.CodeUnauthorized, errors.New("password is wrong"))
	}

	return &dto.SigninResponse{ID: user.ID.String(), Email: user.Email}, nil
}

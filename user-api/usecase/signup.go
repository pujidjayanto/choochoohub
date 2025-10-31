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

type SignUpUsecase interface {
	Create(c context.Context, req dto.SignupRequest) error
}

type signUpUsecase struct {
	userRepository repository.UserRepository
	eventBus       eventbus.EventBus
}

func NewSignupUsecase(userRepository repository.UserRepository, eventBus eventbus.EventBus) SignUpUsecase {
	return &signUpUsecase{userRepository, eventBus}
}

func (signUpUsecase *signUpUsecase) Create(c context.Context, req dto.SignupRequest) error {
	hashedPwd, err := stringhash.Hash(req.Password)
	if err != nil {
		return apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
	}

	user, err := signUpUsecase.userRepository.Create(c, &model.User{PasswordHash: hashedPwd, Email: req.Email})
	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return apperror.NewAppError(http.StatusUnprocessableEntity, apperror.CodeValidationFailed, err)
		}
		return apperror.NewAppError(http.StatusInternalServerError, apperror.CodeInternalServerError, err)
	}

	signUpUsecase.eventBus.Publish("user.created", dto.OtpRequest{
		UserId:      user.ID,
		Channel:     string(model.UserOtpChannelEmail),
		Destination: user.Email,
		Purpose:     string(model.UserOtpPurposeSignup),
		ExpiredAt:   time.Now().Add(5 * time.Minute),
	})

	return nil
}

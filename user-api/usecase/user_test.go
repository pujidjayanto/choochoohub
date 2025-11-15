package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/matryer/is"
	"gorm.io/gorm"

	"github.com/pujidjayanto/choochoohub/user-api/apperror"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/mocks"
	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/stringhash"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

func TestUserUsecase_Signup(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		is := is.New(t)

		req := dto.SignupRequest{
			Email:    "test@example.com",
			Password: "secure123",
		}

		userID := uuid.New()
		mockRepo := &mocks.UserRepositoryMock{
			CreateFunc: func(_ context.Context, user *model.User) (*model.User, error) {
				is.True(user.PasswordHash != req.Password)
				is.Equal(user.Email, req.Email)
				return &model.User{
					ID:           userID,
					Email:        user.Email,
					PasswordHash: user.PasswordHash,
				}, nil
			},
		}

		eventCalled := false
		mockBus := &mocks.EventBusMock{
			PublishFunc: func(event string, payload any) {
				is.Equal(event, "user.created")

				otpReq, ok := payload.(dto.OtpRequest)
				is.True(ok)
				is.Equal(otpReq.UserId, userID)
				is.Equal(otpReq.Channel, string(model.UserOtpChannelEmail))
				is.Equal(otpReq.Destination, "test@example.com")
				is.Equal(otpReq.Purpose, string(model.UserOtpPurposeSignup))
				is.True(otpReq.ExpiredAt.After(time.Now()))

				eventCalled = true
			},
		}

		uc := usecase.NewUserUsecase(mockRepo, mockBus)
		err := uc.Signup(ctx, req)

		is.NoErr(err)
		is.True(eventCalled)
	})

	t.Run("duplicate email", func(t *testing.T) {
		is := is.New(t)

		req := dto.SignupRequest{
			Email:    "exists@example.com",
			Password: "secure123",
		}

		mockRepo := &mocks.UserRepositoryMock{
			CreateFunc: func(_ context.Context, user *model.User) (*model.User, error) {
				return nil, gorm.ErrDuplicatedKey
			},
		}
		mockBus := &mocks.EventBusMock{}

		uc := usecase.NewUserUsecase(mockRepo, mockBus)
		err := uc.Signup(ctx, req)

		appErr, ok := err.(*apperror.AppError)
		is.True(ok)
		is.Equal(appErr.StatusCode, http.StatusUnprocessableEntity)
		is.Equal(appErr.Err, gorm.ErrDuplicatedKey)
	})

	t.Run("unexpected repository error", func(t *testing.T) {
		is := is.New(t)

		mockRepo := &mocks.UserRepositoryMock{
			CreateFunc: func(_ context.Context, user *model.User) (*model.User, error) {
				return nil, errors.New("db fail")
			},
		}

		uc := usecase.NewUserUsecase(mockRepo, &mocks.EventBusMock{})

		err := uc.Signup(ctx, dto.SignupRequest{
			Email:    "x@example.com",
			Password: "123",
		})

		appErr, ok := err.(*apperror.AppError)
		is.True(ok)
		is.Equal(appErr.StatusCode, http.StatusInternalServerError)
		is.Equal(appErr.Err.Error(), "db fail")
	})
}

func TestUserUsecase_Signin(t *testing.T) {
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		is := is.New(t)

		hashed, _ := stringhash.Hash("password123")
		userID := uuid.New()

		mockRepo := &mocks.UserRepositoryMock{
			FindByEmailFunc: func(_ context.Context, email string) (*model.User, error) {
				is.Equal(email, "user@example.com")
				return &model.User{
					ID:           userID,
					Email:        email,
					PasswordHash: hashed,
				}, nil
			},
		}

		uc := usecase.NewUserUsecase(mockRepo, &mocks.EventBusMock{})

		resp, err := uc.Signin(ctx, dto.SigninRequest{
			Email:    "user@example.com",
			Password: "password123",
		})

		is.NoErr(err)
		is.Equal(resp.ID, userID.String())
		is.Equal(resp.Email, "user@example.com")
	})

	t.Run("user not found", func(t *testing.T) {
		is := is.New(t)

		mockRepo := &mocks.UserRepositoryMock{
			FindByEmailFunc: func(_ context.Context, email string) (*model.User, error) {
				return nil, errors.New("not found")
			},
		}

		uc := usecase.NewUserUsecase(mockRepo, &mocks.EventBusMock{})

		_, err := uc.Signin(ctx, dto.SigninRequest{
			Email:    "missing@example.com",
			Password: "xxx",
		})

		appErr, ok := err.(*apperror.AppError)
		is.True(ok)
		is.Equal(appErr.StatusCode, http.StatusUnauthorized)
	})

	t.Run("wrong password", func(t *testing.T) {
		is := is.New(t)

		hashed, _ := stringhash.Hash("correct")

		mockRepo := &mocks.UserRepositoryMock{
			FindByEmailFunc: func(_ context.Context, email string) (*model.User, error) {
				return &model.User{
					ID:           uuid.New(),
					Email:        email,
					PasswordHash: hashed,
				}, nil
			},
		}

		uc := usecase.NewUserUsecase(mockRepo, &mocks.EventBusMock{})

		_, err := uc.Signin(ctx, dto.SigninRequest{
			Email:    "user@example.com",
			Password: "wrong",
		})

		appErr, ok := err.(*apperror.AppError)
		is.True(ok)
		is.Equal(appErr.StatusCode, http.StatusUnauthorized)
	})
}

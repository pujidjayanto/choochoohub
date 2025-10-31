package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/matryer/is"

	"github.com/pujidjayanto/choochoohub/user-api/apperror"
	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/mocks"
	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

func TestSignUpUsecase_Create(t *testing.T) {
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
				is.True(user.PasswordHash != "secure123")
				is.Equal(user.Email, "test@example.com")
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

		signupUC := usecase.NewSignupUsecase(mockRepo, mockBus)
		err := signupUC.Create(ctx, req)
		is.NoErr(err)
		is.True(eventCalled)
	})

	t.Run("repository error", func(t *testing.T) {
		is := is.New(t)

		req := dto.SignupRequest{
			Email:    "exists@example.com",
			Password: "secure123",
		}

		mockRepo := &mocks.UserRepositoryMock{
			CreateFunc: func(_ context.Context, user *model.User) (*model.User, error) {
				return nil, errors.New("duplicate email")
			},
		}

		mockBus := &mocks.EventBusMock{}

		signupUC := usecase.NewSignupUsecase(mockRepo, mockBus)
		err := signupUC.Create(ctx, req)

		appErr, ok := err.(*apperror.AppError)
		is.True(ok)
		is.Equal(appErr.Err.Error(), "duplicate email")
		is.Equal(appErr.StatusCode, http.StatusInternalServerError)
	})
}

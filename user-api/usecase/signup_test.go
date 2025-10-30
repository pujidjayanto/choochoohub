package usecase_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

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

		mockRepo := &mocks.UserRepositoryMock{
			CreateFunc: func(_ context.Context, user *model.User) error {
				is.True(user.PasswordHash != "secure123") // password should be hashed
				is.Equal(user.Email, "test@example.com")
				return nil
			},
		}

		signupUC := usecase.NewSignupUsecase(mockRepo)
		err := signupUC.Create(ctx, req)
		is.NoErr(err)
	})

	t.Run("repository error", func(t *testing.T) {
		is := is.New(t)

		req := dto.SignupRequest{
			Email:    "exists@example.com",
			Password: "secure123",
		}

		mockRepo := &mocks.UserRepositoryMock{
			CreateFunc: func(_ context.Context, user *model.User) error {
				return errors.New("duplicate email")
			},
		}

		signupUC := usecase.NewSignupUsecase(mockRepo)
		err := signupUC.Create(ctx, req)
		appErr, ok := err.(*apperror.AppError)
		is.True(ok)
		is.Equal(appErr.Err.Error(), "duplicate email")
		is.Equal(appErr.StatusCode, http.StatusInternalServerError)
	})
}

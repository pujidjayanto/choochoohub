package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/matryer/is"
	"go.uber.org/mock/gomock"

	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/mocks"
	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

func TestSignUpUsecase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	signupUC := usecase.NewSignupUsecase(mockRepo)
	ctx := context.Background()

	t.Run("success", func(t *testing.T) {
		is := is.New(t)

		req := dto.SignupRequest{
			Email:    "test@example.com",
			Password: "secure123",
		}

		mockRepo.EXPECT().
			Create(ctx, gomock.Any()).
			DoAndReturn(func(_ context.Context, user *model.User) error {
				is.True(user.PasswordHash != "secure123") // password should be hashed
				is.Equal(user.Email, "test@example.com")
				return nil
			})

		err := signupUC.Create(ctx, req)
		is.NoErr(err)
	})

	t.Run("repository error", func(t *testing.T) {
		is := is.New(t)

		req := dto.SignupRequest{
			Email:    "exists@example.com",
			Password: "secure123",
		}

		mockRepo.EXPECT().
			Create(ctx, gomock.Any()).
			Return(errors.New("duplicate email"))

		err := signupUC.Create(ctx, req)
		is.Equal(err.Error(), "duplicate email")
	})
}

package usecase_test

import (
	"context"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/matryer/is"
	"go.uber.org/mock/gomock"

	"github.com/pujidjayanto/choochoohub/user-api/dto"
	"github.com/pujidjayanto/choochoohub/user-api/mocks"
	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/pwd"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

func TestSignInUsecase_SignIn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockUserRepository(ctrl)
	signinUC := usecase.NewSignInUsecase(mockRepo)
	ctx := context.Background()

	hashedPwd, _ := pwd.Hash("secret123")
	userID := uuid.New()

	t.Run("success", func(t *testing.T) {
		is := is.New(t)

		mockRepo.EXPECT().
			FindByEmail(ctx, "test@example.com").
			Return(&model.User{
				ID:           userID,
				Email:        "test@example.com",
				PasswordHash: hashedPwd,
			}, nil)

		resp, err := signinUC.SignIn(ctx, dto.SigninRequest{
			Email:    "test@example.com",
			Password: "secret123",
		})

		is.NoErr(err)                            // no error
		is.Equal(resp.Email, "test@example.com") // email matches
		is.Equal(resp.ID, userID.String())       // ID matches
	})

	t.Run("invalid password", func(t *testing.T) {
		is := is.New(t)

		mockRepo.EXPECT().
			FindByEmail(ctx, "test@example.com").
			Return(&model.User{
				ID:           userID,
				Email:        "test@example.com",
				PasswordHash: hashedPwd,
			}, nil)

		resp, err := signinUC.SignIn(ctx, dto.SigninRequest{
			Email:    "test@example.com",
			Password: "wrongpassword",
		})

		is.True(resp == nil)                      // response is nil
		is.Equal(err.Error(), "invalid password") // error message matches
	})

	t.Run("user not found", func(t *testing.T) {
		is := is.New(t)

		mockRepo.EXPECT().
			FindByEmail(ctx, "unknown@example.com").
			Return(nil, errors.New("user not found"))

		resp, err := signinUC.SignIn(ctx, dto.SigninRequest{
			Email:    "unknown@example.com",
			Password: "secret123",
		})

		is.True(resp == nil)                    // response is nil
		is.Equal(err.Error(), "user not found") // error message matches
	})
}

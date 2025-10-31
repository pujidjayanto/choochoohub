package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/matryer/is"

	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
	"github.com/pujidjayanto/choochoohub/user-api/repository/testutils"
)

func TestUserOtpRepository_Create(t *testing.T) {
	is := is.New(t)
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := repository.NewUserOtpRepository(db)

		user := &model.User{
			Email:        "otpuser@example.com",
			PasswordHash: "hashedpassword",
		}

		userRepo := repository.NewUserRepository(db)
		_, err := userRepo.Create(ctx, user)
		is.NoErr(err)
		is.True(user.ID != uuid.Nil)

		otp, err := repo.Create(ctx, &model.UserOtp{
			UserID:      user.ID,
			Channel:     "email",
			Destination: "otpuser@example.com",
			OTPHash:     "hashedotp",
			Purpose:     "signup",
			ExpiresAt:   time.Now().Add(5 * time.Minute),
		})
		is.NoErr(err)

		var fetched model.UserOtp
		err = db.GetDB(ctx).First(&fetched, "id = ?", otp.ID).Error
		is.NoErr(err)
		is.Equal(fetched.UserID, otp.UserID)
		is.Equal(fetched.Channel, otp.Channel)
		is.Equal(fetched.Purpose, otp.Purpose)
		is.Equal(fetched.Status, model.UserOtpStatus("pending"))
	})
}

package repository

import (
	"context"

	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/db"
)

type UserOtpRepository interface {
	Create(ctx context.Context, user *model.UserOtp) (*model.UserOtp, error)
}

type userOtpRepository struct {
	db db.DatabaseHandler
}

func NewUserOtpRepository(db db.DatabaseHandler) UserOtpRepository {
	return &userOtpRepository{db: db}
}

func (r *userOtpRepository) Create(ctx context.Context, otp *model.UserOtp) (*model.UserOtp, error) {
	if err := r.db.GetDB(ctx).Create(otp).Error; err != nil {
		return nil, err
	}

	return otp, nil
}

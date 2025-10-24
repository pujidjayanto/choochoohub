package repository

import (
	"context"

	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/db"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	db db.DatabaseHandler
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.GetDB(ctx).Create(user).Error
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	err := r.db.GetDB(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUserRepository(db db.DatabaseHandler) UserRepository {
	return &userRepository{db: db}
}

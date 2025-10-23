package repository

import (
	"context"

	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/db"
)

type UserRepository interface {
	Create(context.Context, *model.User) error
}

type userRepository struct {
	db db.DatabaseHandler
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	return r.db.GetDB(ctx).Create(user).Error
}

func NewUserRepository(db db.DatabaseHandler) UserRepository {
	return &userRepository{db: db}
}

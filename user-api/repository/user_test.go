package repository_test

import (
	"context"
	"testing"

	"github.com/matryer/is"

	"github.com/pujidjayanto/choochoohub/user-api/model"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
	"github.com/pujidjayanto/choochoohub/user-api/repository/testutils"
)

func TestCreate(t *testing.T) {
	is := is.New(t)
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := repository.NewUserRepository(db)

		user := &model.User{
			Email:        "john@gmail.com",
			PasswordHash: "dummyhashedpassword",
		}

		newUssr, err := repo.Create(ctx, user)
		is.NoErr(err)
		is.Equal(newUssr.Email, "john@gmail.com")
	})
}

func TestFindByEmail(t *testing.T) {
	is := is.New(t)
	db := testutils.NewTestDb(t)

	testutils.WithTransaction(t, db, func(ctx context.Context) {
		repo := repository.NewUserRepository(db)

		user := &model.User{
			Email:        "john@gmail.com",
			PasswordHash: "dummyhashedpassword",
		}

		_, err := repo.Create(ctx, user)
		is.NoErr(err)

		existingUser, err := repo.FindByEmail(ctx, user.Email)
		is.NoErr(err)
		is.Equal(existingUser.Email, "john@gmail.com")
	})
}

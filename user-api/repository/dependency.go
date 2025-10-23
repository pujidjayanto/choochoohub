package repository

import (
	"github.com/pujidjayanto/choochoohub/user-api/pkg/db"
)

type Dependency struct {
	UserRepository UserRepository
}

func NewDependency(dbHandler db.DatabaseHandler) Dependency {
	return Dependency{
		UserRepository: NewUserRepository(dbHandler),
	}
}

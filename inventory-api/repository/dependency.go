package repository

import "github.com/pujidjayanto/choochoohub/inventory-api/pkg/db"

type Dependency struct {
	StationRepository StationRepository
}

func NewDependency(db db.DatabaseHandler) Dependency {
	return Dependency{
		StationRepository: NewStationRepository(db),
	}
}

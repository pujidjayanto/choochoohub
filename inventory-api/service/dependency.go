package service

import "github.com/pujidjayanto/choochoohub/inventory-api/repository"

type Dependency struct {
	StationService StationService
}

func NewDependency(repositories repository.Dependency) Dependency {
	return Dependency{
		StationService: *NewStationService(repositories.StationRepository),
	}
}

package service

import (
	"context"

	"github.com/pujidjayanto/choochoohub/inventory-api/proto"

	"github.com/pujidjayanto/choochoohub/inventory-api/repository"
)

type StationService struct {
	repo repository.StationRepository
}

func NewStationService(repo repository.StationRepository) *StationService {
	return &StationService{repo: repo}
}

func (s *StationService) ListStations(ctx context.Context, req *proto.ListStationsRequest) (*proto.ListStationsResponse, error) {
	stations, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var pbStations []*proto.StationItem
	for _, st := range stations {
		pbStations = append(pbStations, &proto.StationItem{
			Id:   st.ID,
			Code: st.Code,
			Name: st.Name,
			City: st.City,
		})
	}

	return &proto.ListStationsResponse{
		Stations: pbStations,
	}, nil
}

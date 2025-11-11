package service

import (
	"context"

	pb "github.com/pujidjayanto/choochoohub/inventory-api/proto"

	"github.com/pujidjayanto/choochoohub/inventory-api/repository"
)

type StationService struct {
	repo repository.StationRepository
	pb.UnimplementedStationServiceServer
}

func NewStationService(repo repository.StationRepository) *StationService {
	return &StationService{repo: repo}
}

func (s *StationService) ListStations(ctx context.Context, req *pb.ListStationsRequest) (*pb.ListStationsResponse, error) {
	stations, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	var pbStations []*pb.StationItem
	for _, st := range stations {
		pbStations = append(pbStations, &pb.StationItem{
			Id:   st.ID,
			Code: st.Code,
			Name: st.Name,
			City: st.City,
		})
	}

	return &pb.ListStationsResponse{
		Stations: pbStations,
	}, nil
}

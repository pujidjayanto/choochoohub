package grpc

import (
	"context"

	"github.com/pujidjayanto/choochoohub/inventory-api/proto"
	"github.com/pujidjayanto/choochoohub/inventory-api/service"
)

type StationGrpcApi interface {
	ListStations(ctx context.Context, req *proto.ListStationsRequest) (*proto.ListStationsResponse, error)
}

type stationGrpcApi struct {
	proto.UnimplementedStationServiceServer
	stationService *service.StationService
}

func NewStationGrpcApi(stationService *service.StationService) StationGrpcApi {
	return &stationGrpcApi{stationService: stationService}
}

func (a *stationGrpcApi) ListStations(ctx context.Context, req *proto.ListStationsRequest) (*proto.ListStationsResponse, error) {
	return a.stationService.ListStations(ctx, req)
}

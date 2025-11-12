package grpc

import (
	"log"

	"github.com/pujidjayanto/choochoohub/inventory-api/proto"
	"github.com/pujidjayanto/choochoohub/inventory-api/service"
	"google.golang.org/grpc"
)

type Dependency struct {
	StationGrpcApi StationGrpcApi
}

func NewDependency(services service.Dependency) Dependency {
	return Dependency{
		StationGrpcApi: NewStationGrpcApi(&services.StationService),
	}
}

// RegisterGrpc registers the concrete gRPC server with grpc.Server
func RegisterGrpc(grpcServer *grpc.Server, api StationGrpcApi) {
	// Type assert to concrete struct
	srv, ok := api.(*stationGrpcApi)
	if !ok {
		log.Fatal("failed to register gRPC: invalid type")
	}

	proto.RegisterStationServiceServer(grpcServer, srv)
}

package bootstrap

import (
	"github.com/pujidjayanto/choochoohub/inventory-api/pkg/db"
	pb "github.com/pujidjayanto/choochoohub/inventory-api/proto"
	"github.com/pujidjayanto/choochoohub/inventory-api/repository"
	"github.com/pujidjayanto/choochoohub/inventory-api/service"
	"google.golang.org/grpc"
)

type Application struct {
	GrpcServer *grpc.Server
	ServerAddr string
}

func NewApplicationServer() (*Application, error) {
	if err := initConfig(); err != nil {
		return nil, err
	}

	db, err := db.InitDatabaseHandler(GetDatabaseDSN())
	if err != nil {
		return nil, err
	}

	repositories := repository.NewDependency(db)
	services := service.NewDependency(repositories)

	grpcServer := grpc.NewServer()
	pb.RegisterStationServiceServer(grpcServer, &services.StationService)

	return &Application{
		GrpcServer: grpcServer,
		ServerAddr: GetServerPort(),
	}, nil
}

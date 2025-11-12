package bootstrap

import (
	grpcApi "github.com/pujidjayanto/choochoohub/inventory-api/api/grpc"
	"github.com/pujidjayanto/choochoohub/inventory-api/pkg/db"
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
	apis := grpcApi.NewDependency(services)

	grpcServer := grpc.NewServer()
	grpcApi.RegisterGrpc(grpcServer, apis.StationGrpcApi)

	return &Application{
		GrpcServer: grpcServer,
		ServerAddr: GetServerPort(),
	}, nil
}

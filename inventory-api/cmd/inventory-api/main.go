package main

import (
	"net"

	"github.com/pujidjayanto/choochoohub/inventory-api/bootstrap"
	"github.com/pujidjayanto/choochoohub/inventory-api/pkg/logger"
)

func main() {
	log := logger.GetLogger()

	app, err := bootstrap.NewApplicationServer()
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	lis, err := net.Listen("tcp", app.ServerAddr)
	if err != nil {
		log.Fatalf("failed to initialize application: %v", err)
	}

	if err := app.GrpcServer.Serve(lis); err != nil {
		log.Fatalf("grpc server failed: %v", err)
	}
}

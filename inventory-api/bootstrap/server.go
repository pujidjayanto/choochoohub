package bootstrap

import "google.golang.org/grpc"

func NewApplicationServer() (*grpc.Server, error) {
	if err := initConfig(); err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer()

	return grpcServer, nil
}

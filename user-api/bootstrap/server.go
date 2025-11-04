package bootstrap

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/api"
	"github.com/pujidjayanto/choochoohub/user-api/pkg"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
)

func NewApplicationServer() (*http.Server, CleanupFunc, error) {
	if err := initConfig(); err != nil {
		return nil, nil, err
	}

	sharedDependency, err := pkg.NewDependency(GetDatabaseDSN(), GetKafkaHost())
	if err != nil {
		return nil, nil, err
	}

	repositories := repository.NewDependency(sharedDependency.DB)
	usecases := usecase.NewDependency(repositories, sharedDependency.EventBus)
	apis := api.NewDependency(usecases)

	RegisterOtpSubscriber(sharedDependency, usecases.OtpUsecase)
	VerifiedOtpSubscriber(sharedDependency, repositories.UserRepository)

	router := echo.New()
	routes(router, apis, sharedDependency.Logger)

	server := &http.Server{
		Addr:         GetServerPort(),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	httpCleanup := CleanupFunc(func(ctx context.Context) error {
		return server.Shutdown(ctx)
	})

	dbCleanup := CleanupFunc(func(ctx context.Context) error {
		return sharedDependency.DB.Close()
	})

	cleanup := ChainCleanup(httpCleanup, dbCleanup)

	return server, cleanup, nil
}

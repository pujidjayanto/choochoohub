package bootstrap

import (
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/api"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/db"
	"github.com/pujidjayanto/choochoohub/user-api/repository"
	"github.com/pujidjayanto/choochoohub/user-api/usecase"
	"gorm.io/gorm"
)

func NewApplicationServer() (*http.Server, CleanupFunc, error) {
	if err := initConfig(); err != nil {
		return nil, nil, err
	}

	database, err := db.InitDatabaseHandler(GetDatabaseDSN(), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		TranslateError:         true,
		NowFunc:                func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return nil, nil, err
	}

	if err = database.Ping(context.Background()); err != nil {
		return nil, nil, err
	}

	repositories := repository.NewDependency(database)
	usecases := usecase.NewDependency(repositories)
	apis := api.NewDependency(usecases)

	router := echo.New()
	routes(router, apis)

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
		return database.Close()
	})

	cleanup := ChainCleanup(httpCleanup, dbCleanup)

	return server, cleanup, nil
}

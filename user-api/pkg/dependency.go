package pkg

import (
	"context"
	"time"

	"github.com/pujidjayanto/choochoohub/user-api/pkg/db"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/eventbus"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Dependency struct {
	DB       db.DatabaseHandler
	Logger   *logrus.Logger
	EventBus eventbus.EventBus
}

func NewDependency(dbDsn string) (*Dependency, error) {
	database, err := db.InitDatabaseHandler(dbDsn, &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
		TranslateError:         true,
		NowFunc:                func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return nil, err
	}

	if err = database.Ping(context.Background()); err != nil {
		return nil, err
	}

	return &Dependency{
		DB:       database,
		EventBus: eventbus.New(),
		Logger:   logger.GetLogger(),
	}, nil
}

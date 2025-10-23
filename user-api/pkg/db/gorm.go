package db

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseHandler interface {
	GetDB(ctx context.Context) *gorm.DB
	RunTransaction(ctx context.Context, fc func(ctx context.Context) error) error
	Ping(ctx context.Context) error
	Close() error
}

type databaseHandler struct {
	db  *gorm.DB
	sql *sql.DB
}

type contextKey string

const txKey = contextKey("DBTX")

func InitDatabaseHandler(dsn string, gormConfig *gorm.Config) (DatabaseHandler, error) {
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	dbPool, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error opening database pool: %w", err)
	}

	dbPool.SetMaxIdleConns(10)
	dbPool.SetMaxOpenConns(100)
	dbPool.SetConnMaxLifetime(time.Hour)

	err = dbPool.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't ping the db, %v", err)
	}

	return &databaseHandler{db: db, sql: dbPool}, nil
}

func (h *databaseHandler) GetDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok {
		return tx.WithContext(ctx)
	}
	return h.db.WithContext(ctx)
}

func (h *databaseHandler) Ping(ctx context.Context) error {
	return h.sql.PingContext(ctx)
}

func (h *databaseHandler) RunTransaction(ctx context.Context, fc func(ctx context.Context) error) error {
	// If there's already a transaction in the context, reuse it
	if _, ok := ctx.Value(txKey).(*gorm.DB); ok {
		return fc(ctx)
	}

	// If no transaction exists, create a new one
	return h.db.Transaction(func(tx *gorm.DB) error {
		txCtx := context.WithValue(ctx, txKey, tx)
		return fc(txCtx)
	})
}

func (h *databaseHandler) Close() error {
	err := h.sql.Close()
	if err != nil {
		return err
	}

	return nil
}

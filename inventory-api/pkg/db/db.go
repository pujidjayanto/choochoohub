package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DatabaseHandler interface {
	GetPool(ctx context.Context) *pgxpool.Pool
	RunTransaction(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) error
	Ping(ctx context.Context) error
	Close()
}

type databaseHandler struct {
	pool *pgxpool.Pool
}

type contextKey string

const txKey = contextKey("PGXTX")

func InitDatabaseHandler(dsn string) (DatabaseHandler, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse DB config: %w", err)
	}

	config.MaxConns = 20
	config.MinConns = 2
	config.MaxConnLifetime = time.Hour
	config.HealthCheckPeriod = 30 * time.Second

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("failed to create pgx pool: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	return &databaseHandler{pool: pool}, nil
}

func (h *databaseHandler) GetPool(ctx context.Context) *pgxpool.Pool {
	return h.pool
}

func (h *databaseHandler) Ping(ctx context.Context) error {
	return h.pool.Ping(ctx)
}

// RunTransaction executes fn in a transaction.
// If a transaction already exists in context, it reuses it.
func (h *databaseHandler) RunTransaction(ctx context.Context, fn func(ctx context.Context, tx pgx.Tx) error) error {
	// reuse existing tx if available
	if tx, ok := ctx.Value(txKey).(pgx.Tx); ok {
		return fn(ctx, tx)
	}

	// start a new transaction
	tx, err := h.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("failed to begin tx: %w", err)
	}
	defer func() {
		// rollback if not committed
		_ = tx.Rollback(ctx)
	}()

	txCtx := context.WithValue(ctx, txKey, tx)

	if err := fn(txCtx, tx); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("failed to commit tx: %w", err)
	}

	return nil
}

// Close shuts down the pool
func (h *databaseHandler) Close() {
	h.pool.Close()
}

package testutils

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/db"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/envloader"
	"gorm.io/gorm"
)

func loadTestDatabaseDsn() (string, error) {
	envPath, err := envloader.GetEnvPath()
	if err != nil || envPath == "" {
		return "", fmt.Errorf("no .env file found")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return "", err
	}

	var (
		host     = os.Getenv("TEST_DB_HOST")
		port     = os.Getenv("TEST_DB_PORT")
		user     = os.Getenv("TEST_DB_USER")
		password = os.Getenv("TEST_DB_PASSWORD")
		dbName   = os.Getenv("TEST_DB_NAME")
	)

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName,
	), nil
}

// NewTestDb creates a new test database handler
func NewTestDb(t *testing.T) db.DatabaseHandler {
	dsn, err := loadTestDatabaseDsn()
	if err != nil {
		t.Fatalf("failed to load test database dsn: %v", err)
	}

	database, err := db.InitDatabaseHandler(dsn, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to initialize database: %v", err)
	}

	// Register cleanup to close the database connection
	t.Cleanup(func() {
		//nolint:errcheck
		database.Close()
	})

	return database
}

// WithTransaction runs the test function within a transaction and rolls it back afterward
func WithTransaction(t *testing.T, db db.DatabaseHandler, testFn func(ctx context.Context)) {
	ctx := context.Background()

	err := db.RunTransaction(ctx, func(txCtx context.Context) error {
		testFn(txCtx)
		// Always return an error to force rollback
		return fmt.Errorf("force rollback")
	})

	// We expect an error because we're forcing a rollback
	if err == nil {
		t.Fatal("expected error to force rollback, got nil")
	}

	if err.Error() != "force rollback" {
		t.Fatalf("expected 'force rollback' error, got: %v", err)
	}
}

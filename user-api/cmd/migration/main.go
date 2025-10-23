package main

import (
	"database/sql"
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/logger"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log := logger.GetLogger()

	var (
		doMigration   bool
		doSeed        bool
		migrationName string
		seedName      string
	)

	flag.BoolVar(&doMigration, "migrate", false, "run database migrations")
	flag.BoolVar(&doSeed, "seed", false, "run seeds")
	flag.StringVar(&migrationName, "create-migration", "", "create a new migration file")
	flag.StringVar(&seedName, "create-seed", "", "create a new seed file")
	flag.Parse()

	switch {
	case migrationName != "":
		createNewMigration(migrationName, log)
	case seedName != "":
		createNewSeed(seedName, log)
	case doMigration:
		runMigrationCommand(log)
	case doSeed:
		runSeedCommand(log)
	default:
		log.Info("no command provided, use -migrate, -seed, -create-migration, or -create-seed")
	}
}

func runMigrationCommand(log *logrus.Logger) {
	args := flag.Args()
	if len(args) < 1 {
		log.Fatal("please specify 'up' or 'down' for migration")
	}
	direction := args[0]

	env, err := loadConfiguration()
	if err != nil {
		log.WithError(err).Fatal("failed to load configuration")
	}

	db, err := prepareDatabase(env.databseDsn())
	if err != nil {
		log.WithError(err).Fatal("failed to prepare database")
	}

	runMigrations(db, direction, env.Migration.Env, log)

	// run on test database in development
	if env.Migration.Env == "development" {
		dbTest, err := prepareDatabase(env.testDatabaseDsn())
		if err != nil {
			log.WithField("dsn", env.testDatabaseDsn()).WithError(err).Fatal("failed to prepare test database")
		}
		runMigrations(dbTest, direction, "test", log)
	}
}

func runSeedCommand(log *logrus.Logger) {
	env, err := loadConfiguration()
	if err != nil {
		log.WithError(err).Fatal("failed to load configuration")
	}

	db, err := prepareDatabase(env.databseDsn())
	if err != nil {
		log.WithError(err).Fatal("failed to prepare database")
	}

	runSeeds(db, env.Migration.Env, log)
}

func prepareDatabase(dsn string) (*sql.DB, error) {
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return sqlDB, nil
}

func runMigrations(db *sql.DB, direction, envName string, log *logrus.Logger) {
	migrationDir := filepath.Join("migrations")
	goose.SetTableName("migration_history")

	switch direction {
	case "up":
		if err := goose.Up(db, migrationDir); err != nil {
			log.WithField("dir", migrationDir).WithField("env", envName).WithError(err).Fatal("failed to run migrations")
		}
		log.WithField("dir", migrationDir).WithField("env", envName).Info("migrations applied successfully")
	case "down":
		if err := goose.Down(db, migrationDir); err != nil {
			log.WithField("dir", migrationDir).WithField("env", envName).WithError(err).Fatal("failed to rollback migrations")
		}
		log.WithField("dir", migrationDir).WithField("env", envName).Info("migrations rolled back successfully")
	default:
		log.WithField("direction", direction).WithField("env", envName).Fatal("invalid migration direction, use 'up' or 'down'")
	}
}

func runSeeds(db *sql.DB, envName string, log *logrus.Logger) {
	seedDir := filepath.Join("migrations", "seeds")
	goose.SetTableName("seed_history")

	if err := goose.Up(db, seedDir); err != nil {
		log.WithField("dir", seedDir).WithField("env", envName).WithError(err).Fatal("failed to run seeds")
	}

	log.WithField("dir", seedDir).WithField("env", envName).Info("seeds applied successfully")
}

func createNewMigration(name string, log *logrus.Logger) {
	migrationDir := filepath.Join("migrations")
	if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
		if err := os.MkdirAll(migrationDir, os.ModePerm); err != nil {
			log.WithField("dir", migrationDir).WithError(err).Fatal("failed to create migrations directory")
		}
	}

	if err := goose.Create(nil, migrationDir, name, "sql"); err != nil {
		log.WithField("dir", migrationDir).WithError(err).Fatal("failed to create migration file")
	}

	log.WithField("migration", name).WithField("dir", migrationDir).Info("migration file created successfully")
}

func createNewSeed(name string, log *logrus.Logger) {
	seedDir := filepath.Join("migrations", "seeds")
	if _, err := os.Stat(seedDir); os.IsNotExist(err) {
		if err := os.MkdirAll(seedDir, os.ModePerm); err != nil {
			log.WithField("dir", seedDir).WithError(err).Fatal("failed to create seeds directory")
		}
	}

	if err := goose.Create(nil, seedDir, name, "sql"); err != nil {
		log.WithField("dir", seedDir).WithError(err).Fatal("failed to create seed file")
	}

	log.WithField("seed", name).WithField("dir", seedDir).Info("seed file created successfully")
}

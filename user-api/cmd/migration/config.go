package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/envloader"
)

type database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Ssl      string
}

type migration struct {
	Env string
}

type Config struct {
	Database     database
	TestDatabase database
	Migration    migration
}

func loadConfiguration() (*Config, error) {
	envPath, err := envloader.GetEnvPath()
	if err != nil || strings.TrimSpace(envPath) == "" {
		return nil, fmt.Errorf("no .env file found")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	return &Config{
		Database: database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Ssl:      os.Getenv("DB_SSL_MODE"),
		},
		TestDatabase: database{
			Host:     os.Getenv("TEST_DB_HOST"),
			Port:     os.Getenv("TEST_DB_PORT"),
			User:     os.Getenv("TEST_DB_USER"),
			Password: os.Getenv("TEST_DB_PASSWORD"),
			Name:     os.Getenv("TEST_DB_NAME"),
			Ssl:      os.Getenv("TEST_DB_SSL_MODE"),
		},
		Migration: migration{
			Env: os.Getenv("SERVER_ENV"),
		},
	}, nil
}

func (e *Config) databaseDsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=UTC",
		e.Database.Host,
		e.Database.Port,
		e.Database.User,
		e.Database.Password,
		e.Database.Name,
		e.Database.Ssl,
	)
}

func (e *Config) testDatabaseDsn() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s timezone=UTC",
		e.TestDatabase.Host,
		e.TestDatabase.Port,
		e.TestDatabase.User,
		e.TestDatabase.Password,
		e.TestDatabase.Name,
		e.TestDatabase.Ssl,
	)
}

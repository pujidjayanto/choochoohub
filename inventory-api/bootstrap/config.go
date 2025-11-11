package bootstrap

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/envloader"
)

// keep env config private, only access global variable via utility functions
type database struct {
	host     string
	port     string
	user     string
	password string
	name     string
	ssl      string
}

type server struct {
	port      string
	env       string
	name      string
	secretKey string
}

type redis struct {
	host         string
	port         string
	user         string
	password     string
	databaseName string
}

type kafka struct {
	host string
	port string
}

type config struct {
	database database
	server   server
	redis    redis
	kafka    kafka
}

var (
	globalConfig *config
	loadOnce     sync.Once
	loadError    error
)

func initConfig() error {
	loadOnce.Do(func() {
		loadError = loadConfiguration()
	})
	return loadError
}

func loadConfiguration() error {
	envPath, err := envloader.GetEnvPath()
	if err != nil || strings.TrimSpace(envPath) == "" {
		return fmt.Errorf("no .env file found")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return err
	}

	globalConfig = &config{
		database: database{
			host:     os.Getenv("DB_HOST"),
			port:     os.Getenv("DB_PORT"),
			user:     os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASSWORD"),
			name:     os.Getenv("DB_NAME"),
			ssl:      os.Getenv("DB_SSL_MODE"),
		},
		server: server{
			port:      os.Getenv("SERVER_PORT"),
			env:       os.Getenv("SERVER_ENV"),
			name:      os.Getenv("SERVER_NAME"),
			secretKey: os.Getenv("SERVER_SECRET_KEY"),
		},
		redis: redis{
			host:         os.Getenv("REDIS_HOST"),
			port:         os.Getenv("REDIS_PORT"),
			user:         os.Getenv("REDIS_USER"),
			password:     os.Getenv("REDIS_PASSWORD"),
			databaseName: os.Getenv("REDIS_DB_NAME"),
		},
		kafka: kafka{
			host: os.Getenv("KAFKA_HOST"),
			port: os.Getenv("KAFKA_PORT"),
		},
	}

	return nil
}

// Utility functions
func GetDatabaseDSN() string {
	checkInitialized()
	password := globalConfig.database.password
	if password != "" {
		password = ":" + password
	}

	return fmt.Sprintf(
		"postgres://%s%s@%s:%s/%s?sslmode=%s",
		globalConfig.database.user,
		password,
		globalConfig.database.host,
		globalConfig.database.port,
		globalConfig.database.name,
		globalConfig.database.ssl,
	)
}
func GetServerPort() string {
	checkInitialized()
	port := globalConfig.server.port
	if port == "" {
		port = "3000"
	}
	return fmt.Sprintf(":%s", port)
}

func GetServerEnv() string {
	checkInitialized()
	return globalConfig.server.env
}

func GetSecretKey() string {
	checkInitialized()
	return globalConfig.server.secretKey
}

func GetKafkaHost() string {
	checkInitialized()
	return fmt.Sprintf("%s:%s", globalConfig.kafka.host, globalConfig.kafka.port)
}

func checkInitialized() {
	if globalConfig == nil {
		panic("configuration not initialized. Call Initialize() first")
	}
}

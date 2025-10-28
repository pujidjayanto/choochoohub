package bootstrap

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/envloader"
)

type server struct {
	port      string
	env       string
	name      string
	secretKey string
}

type userApi struct {
	host string
	port string
}

type config struct {
	server  server
	userApi userApi
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
		server: server{
			port:      os.Getenv("SERVER_PORT"),
			env:       os.Getenv("SERVER_ENV"),
			name:      os.Getenv("SERVER_NAME"),
			secretKey: os.Getenv("SERVER_SECRET_KEY"),
		},
		userApi: userApi{
			host: os.Getenv("USER_API_HOST"),
			port: os.Getenv("USER_API_PORT"),
		},
	}

	return nil
}

func GetServerPort() string {
	checkInitialized()
	port := globalConfig.server.port
	if port == "" {
		port = "3000"
	}
	return fmt.Sprintf(":%s", port)
}

func GetUserApiHost() string {
	checkInitialized()
	return globalConfig.userApi.host
}

func GetUserApiPort() string {
	checkInitialized()
	return globalConfig.userApi.port
}

func checkInitialized() {
	if globalConfig == nil {
		panic("configuration not initialized. Call Initialize() first")
	}
}

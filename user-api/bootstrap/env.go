package bootstrap

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	AppEnv           string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseName     string
}

func LoadEnv() error {
	return godotenv.Load()
}

func GetEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

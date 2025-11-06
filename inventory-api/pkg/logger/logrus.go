package logger

import (
	"os"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

var (
	instance *logrus.Logger
	once     sync.Once
)

// GetLogger returns a singleton logger (info level for all environments)
func GetLogger() *logrus.Logger {
	once.Do(func() {
		instance = logrus.New()
		instance.Out = os.Stdout
		instance.Level = logrus.InfoLevel
		instance.Formatter = &logrus.JSONFormatter{
			TimestampFormat: time.RFC3339,
		}
	})
	return instance
}

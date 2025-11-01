package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pujidjayanto/choochoohub/user-api/api"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/validator"
	"github.com/sirupsen/logrus"
)

func routes(router *echo.Echo, apis api.Dependency, log *logrus.Logger) {
	router.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))
	router.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:    true,
		LogURI:       true,
		LogError:     true,
		LogRequestID: true,
		HandleError:  true,
		LogMethod:    true,
		LogLatency:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				log.
					WithField("request_id", v.RequestID).
					WithField("uri", v.URI).
					WithField("status", v.Status).
					WithField("method", v.Method).
					WithField("latency", v.Latency.Seconds()).
					Info("request")
			} else {
				log.
					WithField("uri", v.URI).
					WithField("request_id", v.RequestID).
					WithField("status", v.Status).
					WithField("error", v.Error.Error()).
					WithField("method", v.Method).
					WithField("latency", v.Latency.Seconds()).
					Info("request_error")
			}
			return nil
		},
	}))

	router.Validator = validator.New()
	router.GET("/ping", apis.PingApi.Ping)

	v1 := router.Group("v1")
	v1.POST("/signup", apis.UserApi.SignUp)
	v1.POST("/signup/verify-otp", apis.OtpApi.Verify)
}

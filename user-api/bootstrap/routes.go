package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pujidjayanto/choochoohub/user-api/api"
	"github.com/sirupsen/logrus"
)

func routes(router *echo.Echo, apis api.Dependency, log *logrus.Logger) {
	// todo: i think request id must be from api-gateway rather than generate new
	router.Use(middleware.RequestID())
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
					WithField("latency", v.Latency.Abs()).
					Info("request")
			} else {
				log.
					WithField("uri", v.URI).
					WithField("request_id", v.RequestID).
					WithField("status", v.Status).
					WithField("error", v.Error.Error()).
					WithField("method", v.Method).
					WithField("latency", v.Latency.Abs()).
					Info("request_error")
			}
			return nil
		},
	}))
	router.GET("/ping", apis.PingController.Ping)

	v1 := router.Group("v1")
	v1.POST("/signup", apis.SignUpController.SignUp)
	v1.POST("/signin", apis.SignInController.SignIn)
}

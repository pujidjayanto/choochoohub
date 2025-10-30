package api

import (
	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/pkg/delivery"
)

type PingApi interface {
	Ping(c echo.Context) error
}

type pingApi struct{}

func NewPingApi() PingApi {
	return &pingApi{}
}

func (pingApi *pingApi) Ping(c echo.Context) error {
	return delivery.Success(c, "pong")
}

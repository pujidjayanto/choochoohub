package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type PingController interface {
	Ping(c echo.Context) error
}

type pingController struct{}

func NewPingController() PingController {
	return &pingController{}
}

func (pingController *pingController) Ping(c echo.Context) error {
	return c.String(http.StatusOK, "pong")
}

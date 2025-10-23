package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/api"
)

func routes(router *echo.Echo, apis api.Dependency) {
	router.POST("/signup", apis.SignUpController.SignUp)
}

package bootstrap

import (
	"github.com/labstack/echo/v4"
	"github.com/pujidjayanto/choochoohub/user-api/api"
)

func routes(router *echo.Echo, apis api.Dependency) {
	v1 := router.Group("v1")
	v1.POST("/signup", apis.SignUpController.SignUp)
}

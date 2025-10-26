package bootstrap

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/pujidjayanto/choochoohub/api-gateway/api"
)

func routes(app *fiber.App, apis api.Dependency) {
	app.Use(requestid.New())
	v1 := app.Group("v1")
	v1.Get("/signin")
}

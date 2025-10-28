package bootstrap

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/pujidjayanto/choochoohub/api-gateway/api"
	"github.com/sirupsen/logrus"
)

func routes(app *fiber.App, apis api.Dependency, log *logrus.Logger) {
	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next() // Process request

		duration := time.Since(start)
		requestId := c.Locals("requestid")
		log.WithFields(logrus.Fields{
			"method":     c.Method(),
			"path":       c.Path(),
			"status":     c.Response().StatusCode(),
			"duration":   duration.Seconds(),
			"request_id": requestId,
		}).Info("Incoming request")

		return err
	})

	v1 := app.Group("v1")
	v1.Post("/signin", apis.UserApi.Signin)
	v1.Post("/signup", apis.UserApi.Signup)
}

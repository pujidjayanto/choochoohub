package delivery

import (
	"runtime/debug"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SuccessNoContent(c *fiber.Ctx) {
	c.SendStatus(fiber.StatusNoContent)
}

func SuccessCreated(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusCreated)
}

func Success(c *fiber.Ctx, data any) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Message: SuccessMessage,
		Data:    data,
	})
}

func Failed(c *fiber.Ctx, statusCode int, message string) error {
	r := ErrorResponse{
		Error: message,
	}

	if header := c.Get("X-Enable-Trace"); header != "" {
		if isAppTrace, err := strconv.ParseBool(header); err == nil && isAppTrace {
			r.Trace = string(debug.Stack())
		}
	}

	c.Set("Connection", "close")

	return c.Status(statusCode).JSON(r)
}

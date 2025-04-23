package monitoring

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func Middleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		duration := time.Since(start).Seconds()

		status := c.Response().StatusCode()
		method := c.Method()
		path := c.Route().Path

		HttpRequests.WithLabelValues(method, path, formatStatus(status)).Inc()
		HttpDuration.WithLabelValues(method, path, formatStatus(status)).Observe(duration)

		return err
	}
}

func formatStatus(status int) string {
	return fiber.StatusMessage(status) // Let's implement this ourselves.
}

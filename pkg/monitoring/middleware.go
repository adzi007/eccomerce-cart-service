package monitoring

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
)

func PrometheusMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start).Seconds()
		method := c.Method()
		path := c.Route().Path
		status := strconv.Itoa(c.Response().StatusCode())

		HttpDuration.WithLabelValues(method, path, status).Observe(duration)
		HttpRequests.WithLabelValues(method, path, status).Inc()

		return err
	}
}

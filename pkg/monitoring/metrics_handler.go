package monitoring

import (
	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// MetricsHandler exposes /metrics using a bridge between fasthttp and net/http
func MetricsHandler() fiber.Handler {
	// Create the standard prometheus HTTP handler
	handler := promhttp.Handler()

	// Adapt it using fasthttpadaptor
	return func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandler(handler)(c.Context())
		return nil
	}
}

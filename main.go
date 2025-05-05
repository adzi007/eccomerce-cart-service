package main

import (
	"cart-service/config"
	"cart-service/config/database"
	"cart-service/pkg/logger"
	"cart-service/pkg/monitoring"
	"cart-service/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

//	@title			Ecommerce Cart Service
//	@version		1.0
//	@description	This is a sample swagger for Fiber
//	@termsOfService	http://swagger.io/terms/

func main() {

	config.LoadConfig()

	mylog := logger.NewLogger()

	// db := database.NewMysqlDatabase()
	var db database.Database

	// fmt.Println("DB_DRIVER >>>>> ", config.ENV.DB_DRIVER)

	if config.ENV.DB_DRIVER == "sqlite" {
		db = database.NewSqliteDatabase()
	} else {
		db = database.NewMysqlDatabase()
	}

	servernya := server.NewFiberServer(db)

	servernya.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &mylog,
	}))

	servernya.Use(limiter.New(limiter.Config{
		Max:        10,               // 10 requests
		Expiration: 30 * time.Second, // per 30 seconds
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP() // limit per IP
		},
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
		},
	}))

	// Register Prometheus metrics
	monitoring.RegisterMetrics()

	// Add middleware
	servernya.Use(monitoring.PrometheusMiddleware())

	grpcServer := server.NewGrpcServer(db)

	go grpcServer.StartGRPCServer()

	go servernya.Start()

	// Graceful shutdown handling
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down servers...")

	// Perform any cleanup if needed
	time.Sleep(1 * time.Second)
	log.Println("Servers stopped.")

}

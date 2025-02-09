package main

import (
	"cart-service/config"
	"cart-service/config/database"
	"cart-service/pkg/logger"
	"cart-service/server"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/contrib/fiberzerolog"
)

//	@title			Ecommerce Cart Service
//	@version		1.0
//	@description	This is a sample swagger for Fiber
//	@termsOfService	http://swagger.io/terms/

func main() {

	config.LoadConfig()

	mylog := logger.NewLogger()

	db := database.NewMysqlDatabase()

	servernya := server.NewFiberServer(db)

	servernya.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &mylog,
	}))

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

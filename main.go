package main

import (
	"cart-service/config"
	"cart-service/config/database"
	"cart-service/pkg/logger"
	"cart-service/server"

	"github.com/gofiber/contrib/fiberzerolog"
)

func main() {

	config.LoadConfig()

	mylog := logger.NewLogger()

	db := database.NewMysqlDatabase()

	servernya := server.NewFiberServer(db)

	servernya.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &mylog,
	}))

	servernya.Start()

	// logger.Error().Err(err).Msg("This is an error message")

}

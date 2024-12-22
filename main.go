package main

import (
	"cart-service/config"
	"cart-service/config/database"
	"cart-service/pkg/logger"
	"cart-service/server"

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

	servernya.Start()

	// logger.Error().Err(err).Msg("This is an error message")

}

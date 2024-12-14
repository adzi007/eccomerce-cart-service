package server

import (
	"cart-service/config/database"
	"cart-service/internal/handler"
	"cart-service/internal/repository"
	"cart-service/internal/usecase"
	"cart-service/pkg/logger"
	"log"

	"github.com/gofiber/fiber/v2"
)

type fiberServer struct {
	app *fiber.App
	db  database.Database
	// conf *config.Config
}

func NewFiberServer(db database.Database) Server {
	fiberApp := fiber.New()
	// fiberApp.Logger.SetLevel(log.DEBUG)

	return &fiberServer{
		app: fiberApp,
		db:  db,
	}
}

func (s *fiberServer) Use(args interface{}) {

	s.app.Use(args)

}

func (s *fiberServer) Start() {
	// Define routes

	s.app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello from Fiber!")
	})

	logger.Info().Msg("This is an info message")
	// logger.Warn().Str("user", "john_doe").Msg("This is a warning message")
	logger.Warn().Msg("This is a warning message")

	s.initializeCartServiceHttpHandler()

	log.Fatal(s.app.Listen(":5000"))
}

func (s *fiberServer) initializeCartServiceHttpHandler() {

	// repository
	cartRepo := repository.NewCartRepository(s.db)

	// use case
	cartUsecase := usecase.NewCartUsecaseImpl(cartRepo)

	// handler

	cartHandler := handler.NewCartHttpHandle(cartUsecase)

	// router
	s.app.Post("/cart", cartHandler.InsertNewCart)

}

package server

import (
	"cart-service/config"
	"cart-service/config/database"
	"cart-service/internal/handler"
	"cart-service/internal/repository"
	productservicerepo "cart-service/internal/repository/product_service_repo"
	"cart-service/internal/usecase"
	"cart-service/pkg/cachestore"
	"cart-service/pkg/logger"
	"context"
	"fmt"
	"log"

	_ "cart-service/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger" // swagger handler
)

type fiberServer struct {
	app *fiber.App
	db  database.Database
	// conf *config.Config
}

func NewFiberServer(db database.Database) Server {
	fiberApp := fiber.New()
	// fiberApp.Logger.SetLevel(log.DEBUG)

	fiberApp.Get("/docs/*", swagger.HandlerDefault)

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

	// s.app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.Status(200).SendString("Hello from Fiber! ini pesan dari admin")
	// })

	logger.Info().Msg("This is an info message")
	// logger.Warn().Str("user", "john_doe").Msg("This is a warning message")
	logger.Warn().Msg("This is a warning message")

	s.initializeCartServiceHttpHandler()

	log.Fatal(s.app.Listen(":5000"))
}

func (s *fiberServer) initializeCartServiceHttpHandler() {

	ctx := context.Background()

	redisHost := config.ENV.REDIS_HOST
	redisPort := config.ENV.REDIS_PORT

	// connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)

	redisConnection := fmt.Sprintf("%s:%s", redisHost, redisPort)

	// redisRepo := cachestore.NewRedisCache(ctx, "ecommerce-redis:6379", "", 0)
	redisRepo := cachestore.NewRedisCache(ctx, redisConnection, "", 0)

	// repository
	cartRepo := repository.NewCartRepository(s.db)

	// product service repository

	productServiceRepo := productservicerepo.NewProductServiceRepository()

	// use case
	cartUsecase := usecase.NewCartUsecaseImpl(cartRepo, redisRepo, productServiceRepo)

	// handler

	cartHandler := handler.NewCartHttpHandle(cartUsecase)

	// lis, err := net.Listen("tcp", ":9001")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// grpcServer := grpc.NewServer()

	// pb.RegisterCartServiceServer(grpcServer, &localGrpc.CartGrpcHandler{})

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %s", err)
	// }

	// router

	// s.app.Get("/metrics", func(c *fiber.Ctx) error {
	// 	promhttp.Handler().ServeHTTP(c.Context().ResponseWriter(), c.Context().Request())
	// 	return nil
	// })

	// s.app.Post("/cart", cartHandler.InsertNewCart)
	s.app.Post("/cart", cartHandler.InsertCart)
	// s.app.Get("/", cartHandler.GetCustomerCart)
	s.app.Get("/:userId", cartHandler.GetCartByCustomer)
	s.app.Put("/", cartHandler.UpdateQty)
	s.app.Delete("/:cartId", cartHandler.DeleteCartItem)
	s.app.Get("/check/redis", cartHandler.Check)

}

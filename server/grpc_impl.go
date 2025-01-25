package server

import (
	pb "cart-service/cart_proto"
	"cart-service/config/database"
	grpchandler "cart-service/internal/handler/grpc_handler"
	"cart-service/internal/repository"
	productservicerepo "cart-service/internal/repository/product_service_repo"
	"cart-service/internal/usecase"
	"cart-service/pkg/cachestore"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type grpcServer struct {
	db database.Database
}

func NewGrpcServer(db database.Database) GrpcServer {

	return &grpcServer{
		db: db,
	}
}

func (s *grpcServer) StartGRPCServer() {

	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	ctx := context.Background()

	redisRepo := cachestore.NewRedisCache(ctx, "localhost:6379", "", 0)

	// repository
	cartRepo := repository.NewCartRepository(s.db)

	// product service repository

	productServiceRepo := productservicerepo.NewProductServiceRepository()

	// use case
	cartUsecase := usecase.NewCartUsecaseImpl(cartRepo, redisRepo, productServiceRepo)

	// grpchandler.NewCartGrpcHandler(cartUsecase)
	grpcCartHandler := grpchandler.NewCartGrpcHandler(cartUsecase)

	pb.RegisterCartServiceServer(grpcServer, grpcCartHandler)
	// pb.RegisterCartServiceServer(grpcServer, grpcCartHadler)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

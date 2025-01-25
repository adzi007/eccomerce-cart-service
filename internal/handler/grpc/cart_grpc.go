package grpc

import (
	pb "cart-service/cart_proto"
	"cart-service/internal/usecase"
	"context"
)

type CartGrpcHandler struct {
	cartUsecase usecase.CartUsecase
	pb.UnimplementedCartServiceServer
}

func NewCartGrpcHandler(usecase usecase.CartUsecase) CartGrpc {
	return &CartGrpcHandler{cartUsecase: usecase}
}

func (h *CartGrpcHandler) GetCartUser(ctx context.Context, req *pb.CartRequest) (*pb.CartResponse, error) {

	dummyData := []*pb.CartItem{
		{
			Id:    1,
			Name:  "Product 1",
			Slug:  "product-1",
			Price: 100,
			Stock: "10",
			Category: &pb.ProductCategory{
				Name: "Category 1",
				Slug: "category-1",
			},
		},
		{
			Id:    2,
			Name:  "Product 2",
			Slug:  "product-2",
			Price: 200,
			Stock: "20",
			Category: &pb.ProductCategory{
				Name: "Category 2",
				Slug: "category-2",
			},
		},
	}

	return &pb.CartResponse{
		Data: dummyData,
	}, nil

	// return nil, nil
}

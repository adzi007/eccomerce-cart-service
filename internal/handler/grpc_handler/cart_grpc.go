package grpchandler

import (
	pb "cart-service/cart_proto"
	"cart-service/internal/usecase"
	"cart-service/pkg/logger"
	"context"
	"fmt"
)

type CartGrpcHandler struct {
	cartUsecase usecase.CartUsecase
	pb.UnimplementedCartServiceServer
}

func NewCartGrpcHandler(usecase usecase.CartUsecase) *CartGrpcHandler {
	return &CartGrpcHandler{cartUsecase: usecase}
}

func (h *CartGrpcHandler) GetCartUser(ctx context.Context, req *pb.CartRequest) (*pb.CartResponse, error) {

	if req == nil {
		logger.Error().Msg("Request is nil")
		return nil, fmt.Errorf("invalid request")
	}

	fmt.Println("req >>>> ", req)

	if h.cartUsecase == nil {
		logger.Error().Msg("CartUsecase is not initialized")
		return nil, fmt.Errorf("internal server error")
	}

	data, err := h.cartUsecase.GetCartByCustomer(req.Id)

	if err != nil {

		logger.Error().Err(err).Msg("Failed to get cart user by ID")

		return nil, err
	}

	var cartData []*pb.CartItem

	for _, v := range data {

		cartItem := &pb.CartItem{
			Id:    uint64(v.ID),
			Name:  v.Name,
			Slug:  v.Slug,
			Price: uint64(v.Price),
			Stock: uint64(v.Stock),
			Category: &pb.ProductCategory{
				Name: v.Category.Name,
				Slug: v.Category.Slug,
			},
		}

		cartData = append(cartData, cartItem)

	}

	return &pb.CartResponse{
		Data: cartData,
	}, nil

}

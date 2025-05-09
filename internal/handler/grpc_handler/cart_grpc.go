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

	if h.cartUsecase == nil {
		logger.Error().Msg("CartUsecase is not initialized")
		return nil, fmt.Errorf("internal server error")
	}

	data, err := h.cartUsecase.GetCartByCustomer(req.Id)

	// pp.Println("data >>> ", data)

	if err != nil {

		logger.Error().Err(err).Msg("Failed to get cart user by ID")

		return nil, err
	}

	var cartData []*pb.CartItem

	for _, v := range data {

		cartItem := &pb.CartItem{
			Id:        uint64(v.ID),
			ProductId: uint64(v.ProductId),
			Name:      v.Name,
			Slug:      v.Slug,
			Price:     uint64(v.Price),
			Qty:       uint64(v.Qty),
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

func (h *CartGrpcHandler) DeleteCartUser(ctx context.Context, req *pb.UserRequest) (*pb.DeleteCartResponse, error) {

	if req == nil {
		logger.Error().Msg("Request is nil")
		return nil, fmt.Errorf("invalid request")
	}

	if h.cartUsecase == nil {
		logger.Error().Msg("CartUsecase is not initialized")
		return nil, fmt.Errorf("internal server error")
	}

	err := h.cartUsecase.DeleteCartByUser(req.UserId)

	if err != nil {

		logger.Error().Err(err).Msg("Failed to get cart user by ID " + req.UserId)

		// return nil, err

		return &pb.DeleteCartResponse{
			Message: "Failed delete cart by user " + req.UserId,
		}, err
	}

	return &pb.DeleteCartResponse{
		Message: "Succes delete cart",
	}, nil

}

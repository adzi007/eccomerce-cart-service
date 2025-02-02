package grpchandler

import (
	proto "cart-service/cart_proto"
	"context"
)

type CartGrpc interface {
	GetCartUser(ctx context.Context, req *proto.CartRequest) (*proto.CartResponse, error)
	DeleteCartUser(ctx context.Context, req *proto.UserRequest) (*proto.DeleteCartResponse, error)
}

package grpc

import (
	proto "cart-service/cart_proto"
	"context"
)

type CartGrpc interface {
	GetCartUser(ctx context.Context, req *proto.CartRequest) (*proto.CartResponse, error)
}

package usecase

import (
	"cart-service/internal/domain"
	"cart-service/internal/model/entity"
)

type CartUsecase interface {
	InsertCart(in *entity.InsertCartDto) error
	CreateNewCart(in *entity.InsertCartDto) error
	GetCustomerCart() error
	GetCartByCustomer(userId string) ([]domain.ProductServiceResponse, error)
	UpdateQty(cartId uint, qty uint) error
	DeleteCartItem(artId uint) error
	Check() error
}

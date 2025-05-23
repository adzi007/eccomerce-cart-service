package usecase

import (
	"cart-service/internal/domain"
	"cart-service/internal/model/entity"
)

type CartUsecase interface {
	InsertCart(in *entity.InsertCartDto) error
	CreateNewCart(in *entity.InsertCartDto) error
	GetCustomerCart() error
	GetCartByCustomer(userId string) ([]domain.ProductCart, error)
	UpdateQty(cartId uint, qty uint) error
	DeleteCartItem(artId uint) error
	DeleteCartByUser(userId string) error
	Check() error
}

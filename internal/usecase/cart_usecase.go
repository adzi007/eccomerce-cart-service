package usecase

import "cart-service/internal/model/entity"

type CartUsecase interface {
	InsertCart(in *entity.InsertCartDto) error
	CreateNewCart(in *entity.InsertCartDto) error
	GetCustomerCart() error
}

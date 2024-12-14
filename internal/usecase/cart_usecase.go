package usecase

import "cart-service/internal/model/entity"

type CartUsecase interface {
	CreateNewCart(in *entity.InsertCartDto) error
}

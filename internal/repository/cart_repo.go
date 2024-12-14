package repository

import "cart-service/internal/model/entity"

type CartRepository interface {
	CreateNewCart(data *entity.InsertCartDto) error
}

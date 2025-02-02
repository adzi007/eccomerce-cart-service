package repository

import "cart-service/internal/model/entity"

type CartRepository interface {
	CreateNewCart(data *entity.InsertCartDto) error
	InsertCart(data *entity.InsertCartDto) error
	GetCartByUser(userId string) (error, []entity.Cart)
	UpdateQty(cartId uint, qty uint) error
	DeleteCartItem(artId uint) error
	DeleteCartByUser(userId string) error
}

package entity

import (
	"time"
)

type InsertCartDto struct {
	UserId    string `json:"user_id" validate:"required"`
	ProductId uint   `json:"product_id" validate:"required"`
	Price     uint   `json:"price" validate:"required"`
	Qty       uint   `json:"qty" validate:"required"`
}

type UpdateCartQtyDto struct {
	ID  uint `json:"cartId" validate:"required"`
	Qty uint `json:"qty" validate:"required"`
}

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserId    string
	ProductId uint
	Price     uint
	Qty       uint
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

// type CartItems struct {
// 	ID        uint `gorm:"primaryKey"`
// 	CartId    uint
// 	ProductId uint
// 	Price     uint
// 	Qty       uint
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// }

package entity

import (
	"time"

	"gorm.io/gorm"
)

type InsertCartDto struct {
	UserId string `json:"user_id"`
}

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserId    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type CartItems struct {
	ID        uint `gorm:"primaryKey"`
	CartId    uint
	ProductId uint
	Price     uint
	Qty       uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

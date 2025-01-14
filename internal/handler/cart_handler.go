package handler

import "github.com/gofiber/fiber/v2"

type CartHandler interface {
	InsertNewCart(ctx *fiber.Ctx) error
	GetCustomerCart(ctx *fiber.Ctx) error
	InsertCart(ctx *fiber.Ctx) error
	GetCartByCustomer(ctx *fiber.Ctx) error
	UpdateQty(ctx *fiber.Ctx) error
	DeleteCartItem(ctx *fiber.Ctx) error
}

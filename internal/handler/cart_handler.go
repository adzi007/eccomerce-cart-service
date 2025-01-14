package handler

import "github.com/gofiber/fiber/v2"

type CartHandler interface {
	InsertNewCart(ctx *fiber.Ctx) error
	GetCustomerCart(ctx *fiber.Ctx) error
	InsertCart(ctx *fiber.Ctx) error
}

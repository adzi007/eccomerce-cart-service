package handler

import (
	"cart-service/internal/model/entity"
	"cart-service/internal/usecase"
	"cart-service/pkg/logger"

	"github.com/gofiber/fiber/v2"
)

type cartHttpHandler struct {
	cartUsecase usecase.CartUsecase
}

func NewCartHttpHandle(cartUc usecase.CartUsecase) CartHandler {
	return &cartHttpHandler{
		cartUsecase: cartUc,
	}
}

func (h *cartHttpHandler) InsertNewCart(ctx *fiber.Ctx) error {

	reqBody := new(entity.InsertCartDto)

	if err := ctx.BodyParser(reqBody); err != nil {
		logger.Error().Err(err).Msg("Error binding request body")

		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})

	}

	if err := h.cartUsecase.CreateNewCart(reqBody); err != nil {

		logger.Error().Err(err).Msg("failed to insert a new cart")

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"pesan": "failed to insert a new cart",
		})

	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "success create a new cart 1",
	})

}

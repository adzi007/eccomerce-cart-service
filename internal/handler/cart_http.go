package handler

import (
	"cart-service/internal/model/entity"
	"cart-service/internal/usecase"
	"cart-service/pkg/logger"
	"strconv"

	"github.com/go-playground/validator/v10"
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

type CartResponse struct {
	Pesan string `json:"pesan"` // Response message
}

// InsertNewCart create a new cart
//
//	@Summary		Create a new cart
//	@Description	API to create a new cart
//	@Tags			Cart
//	@Accept			json
//	@Produce		json
//	@Param			body	body		entity.InsertCartDto	true	"New Cart Details"
//	@Success		200		{object}	handler.CartResponse
//	@Failure		400		{object}	map[string]string
//	@Router			/cart [post]
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

func (h *cartHttpHandler) InsertCart(ctx *fiber.Ctx) error {

	reqBody := new(entity.InsertCartDto)

	if err := ctx.BodyParser(reqBody); err != nil {
		logger.Error().Err(err).Msg("Error binding request body")
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	var validate = validator.New()

	errValidate := validate.Struct(reqBody)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	if err := h.cartUsecase.InsertCart(reqBody); err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"pesan": "failed to insert item to cart",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "success insert item to cart",
	})
}

func (h *cartHttpHandler) GetCustomerCart(ctx *fiber.Ctx) error {

	h.cartUsecase.GetCustomerCart()

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "Test call another service",
	})
}

func (h *cartHttpHandler) GetCartByCustomer(ctx *fiber.Ctx) error {

	userId := ctx.Params("userId")

	data, err := h.cartUsecase.GetCartByCustomer(userId)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"pesan": "Failed to get cart",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "Test get products",
		"data":  data,
	})
}

func (h *cartHttpHandler) UpdateQty(ctx *fiber.Ctx) error {

	reqBody := new(entity.UpdateCartQtyDto)

	if err := ctx.BodyParser(reqBody); err != nil {
		logger.Error().Err(err).Msg("Error binding request body")
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   err.Error(),
		})
	}

	var validate = validator.New()

	errValidate := validate.Struct(reqBody)

	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	if err := h.cartUsecase.UpdateQty(reqBody.ID, reqBody.Qty); err != nil {

		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"pesan": "failed to update item to cart",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "Success Update product",
	})
}

func (h *cartHttpHandler) DeleteCartItem(ctx *fiber.Ctx) error {

	cartId64, errConvert := strconv.ParseUint(ctx.Params("cartId"), 10, 32)

	if errConvert != nil {

		logger.Error().Err(errConvert)
	}

	cartId := uint(cartId64)

	err := h.cartUsecase.DeleteCartItem(cartId)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"pesan": "Failed to delete cart item",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "Success delete cart item",
	})

}

func (h *cartHttpHandler) Check(ctx *fiber.Ctx) error {

	h.cartUsecase.Check()

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"pesan": "Check redis",
	})

}

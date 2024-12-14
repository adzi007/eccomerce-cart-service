package usecase

import (
	"cart-service/internal/model/entity"
	"cart-service/internal/repository"
	"cart-service/pkg/logger"
)

type CartUsecaseImpl struct {
	cartRepo repository.CartRepository
}

func NewCartUsecaseImpl(cartRepo repository.CartRepository) CartUsecase {
	return &CartUsecaseImpl{
		cartRepo: cartRepo,
	}
}

func (c *CartUsecaseImpl) CreateNewCart(in *entity.InsertCartDto) error {

	insertCartdata := &entity.InsertCartDto{
		UserId: in.UserId,
	}

	if err := c.cartRepo.CreateNewCart(insertCartdata); err != nil {

		logger.Error().Err(err).Msg("Failed create new cart")

		return err
	}

	logger.Info().Msg("Success create a new cart 2")

	return nil
}

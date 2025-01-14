package repository

import (
	"cart-service/config/database"
	"cart-service/internal/model/entity"
	"cart-service/pkg/logger"
)

type cartMysqlRepository struct {
	db database.Database
}

func NewCartRepository(db database.Database) CartRepository {
	return &cartMysqlRepository{db: db}
}

func (r *cartMysqlRepository) CreateNewCart(data *entity.InsertCartDto) error {

	cart := &entity.Cart{
		UserId: data.UserId,
	}

	result := r.db.GetDb().Create(cart)

	if result.Error != nil {
		logger.Error().Err(result.Error).Msg("Failed to insert a new cart to the database")
		return result.Error
	}

	logger.Info().Msg("Success insert a new cart to the database")
	return nil

}

func (r *cartMysqlRepository) InsertCart(data *entity.InsertCartDto) error {

	cart := &entity.Cart{
		UserId:    data.UserId,
		ProductId: data.ProductId,
		Price:     data.Price,
		Qty:       data.Qty,
	}

	result := r.db.GetDb().Create(cart)

	if result.Error != nil {
		logger.Error().Err(result.Error).Msg("Failed to insert item to cart")
		return result.Error
	}

	logger.Info().Msg("Success insert a new cart to the database")
	return nil

}

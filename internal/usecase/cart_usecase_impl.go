package usecase

import (
	"cart-service/config"
	"cart-service/internal/model/entity"
	"cart-service/internal/repository"
	"cart-service/pkg/logger"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
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

func (c *CartUsecaseImpl) InsertCart(in *entity.InsertCartDto) error {

	insertCartdata := &entity.InsertCartDto{
		UserId:    in.UserId,
		ProductId: in.ProductId,
		Price:     in.Price,
		Qty:       in.Qty,
	}

	// fmt.Println("insertCartdata >>> ", insertCartdata)

	if err := c.cartRepo.InsertCart(insertCartdata); err != nil {
		logger.Error().Err(err).Msg("Failed insert item to cart")

		return err
	}

	logger.Info().Msg("Success insert item to cart")
	return nil
}

type ApiResponse struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Slug        string   `json:"slug"`
	Description string   `json:"description"`
	PriceBase   int      `json:"priceBase"`
	PriceSell   int      `json:"priceSell"`
	Type        string   `json:"type"`
	Image       string   `json:"image"`
	Stock       int      `json:"stock"`
	Category    Category `json:"category"`
}

type Category struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Todo struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func (c *CartUsecaseImpl) GetCustomerCart() error {

	fmt.Println("coba panggil api service lain", config.ENV.API_GATEWAY)

	url := config.ENV.API_GATEWAY + "/products/products/10"

	request := fiber.Get(url)
	request.Debug()

	_, data, err := request.Bytes()
	if err != nil {
		// panic(err)
		log.Printf("Failed to make request: %v", err)
		// return err
	}

	var apiResponse ApiResponse
	jsonErr := json.Unmarshal(data, &apiResponse)
	if jsonErr != nil {
		// panic(err)
		log.Printf("Failed to to unpack: %v", jsonErr)
	}

	fmt.Println("API Response Data:", apiResponse)

	return nil
}

func (c *CartUsecaseImpl) GetCartByCustomer(userId string) (error, []entity.Cart) {

	err, carts := c.cartRepo.GetCartByUser(userId)

	return err, carts
}

func (c *CartUsecaseImpl) UpdateQty(cartId uint, qty uint) error {

	return c.cartRepo.UpdateQty(cartId, qty)
}

func (c *CartUsecaseImpl) DeleteCartItem(cartId uint) error {

	return c.cartRepo.DeleteCartItem(cartId)
}

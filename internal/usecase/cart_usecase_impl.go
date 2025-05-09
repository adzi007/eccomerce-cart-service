package usecase

import (
	"cart-service/config"
	"cart-service/internal/domain"
	"cart-service/internal/model/entity"
	"cart-service/internal/repository"
	productservicerepo "cart-service/internal/repository/product_service_repo"
	"cart-service/pkg/logger"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/k0kubun/pp/v3"
)

type CartUsecaseImpl struct {
	cartRepo       repository.CartRepository
	cache          domain.CacheRepository
	productService productservicerepo.ProductService
}

func NewCartUsecaseImpl(cartRepo repository.CartRepository, cache domain.CacheRepository, productService productservicerepo.ProductService) CartUsecase {
	return &CartUsecaseImpl{
		cartRepo:       cartRepo,
		cache:          cache,
		productService: productService,
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

	// c.cache.

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

	// url := config.ENV.API_GATEWAY + "/products/products/10"

	request := fiber.Get(config.ENV.API_GATEWAY + "/products/products/10")

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

func (c *CartUsecaseImpl) GetCartByCustomer(userId string) ([]domain.ProductCart, error) {

	// Get cart from internal repository
	err, carts := c.cartRepo.GetCartByUser(userId)
	if err != nil {
		pp.Println("Error fetching carts:", err)
		return nil, err
	}

	if len(carts) == 0 {
		return []domain.ProductCart{}, nil
	}

	var productKeys []string
	productIdsQty := make(map[uint]uint)

	productCartRelation := make(map[uint]uint)

	for _, val := range carts {

		productKeys = append(productKeys, strconv.FormatUint(uint64(val.ProductId), 10))

		productIdsQty[val.ProductId] = val.Qty
		productCartRelation[val.ProductId] = val.ID
	}

	// Check and get product from redis, and get missing producta
	productsFromCache, missingKeyProducts, err := c.cache.MGetProductsCache(productKeys, "product:")
	if err != nil {
		pp.Println("err mget >>> ", err)

		return nil, err
	}

	var productFromService []domain.ProductServiceResponse

	if len(missingKeyProducts) > 0 {

		productCart, err := c.productService.GetProductCart(missingKeyProducts)

		if err != nil {
			pp.Println("err >>", err)

			return nil, err
		}

		productFromService = productCart

		productsToCache := make(map[string]domain.ProductServiceResponse)

		for _, rowProduct := range productCart {

			productkey := strconv.FormatUint(uint64(rowProduct.ID), 10)
			productsToCache["product:"+productkey] = rowProduct

		}

		if err := c.cache.MSetProductsCache(productsToCache, 60); err != nil {
			pp.Println("err mset redis", err)
			return nil, err
		}
	}

	combinedProducts := append(productFromService, productsFromCache...)

	var combinedProductsRelation []domain.ProductCart

	for _, val := range combinedProducts {

		item := domain.ProductCart{
			ID:        productCartRelation[uint(val.ID)],
			ProductId: val.ID,
			Name:      val.Name,
			Slug:      val.Slug,
			Price:     val.Price,
			Qty:       int(productIdsQty[uint(val.ID)]),
			Category:  val.Category,
		}

		combinedProductsRelation = append(combinedProductsRelation, item)

	}

	return combinedProductsRelation, nil
}

func (c *CartUsecaseImpl) UpdateQty(cartId uint, qty uint) error {

	return c.cartRepo.UpdateQty(cartId, qty)
}

func (c *CartUsecaseImpl) DeleteCartItem(cartId uint) error {

	return c.cartRepo.DeleteCartItem(cartId)
}

func (c *CartUsecaseImpl) DeleteCartByUser(userId string) error {

	return c.cartRepo.DeleteCartByUser(userId)
}

func (c *CartUsecaseImpl) Check() error {

	pp.Print("test redissssss xxxx ==============>")

	fmt.Println("test redissssss xxxx ==============>")

	testRedis, err := c.cache.Get("product:5")

	if err != nil {
		pp.Print("error get redis ", err)
		fmt.Println("error get redis >>> ", err)
	}

	pp.Print("testRedis >>> ", testRedis)

	return nil
}

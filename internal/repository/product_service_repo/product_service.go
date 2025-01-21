package productservicerepo

import "cart-service/internal/domain"

// type ProductCartResponse struct {
// 	ID       int    `json:"id"`
// 	Name     string `json:"name"`
// 	Slug     string `json:"slug"`
// 	Price    int    `json:"price"`
// 	Stock    int    `json:"stock"`
// 	Category struct {
// 		Name string `json:"name"`
// 		Slug string `json:"slug"`
// 	} `json:"category"`
// }

type ProductService interface {
	GetProductCart(productIds []uint) ([]domain.ProductCart, error)
}

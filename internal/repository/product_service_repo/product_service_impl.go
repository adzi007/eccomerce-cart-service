package productservicerepo

import (
	"bytes"
	"cart-service/config"
	"cart-service/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type productServiceRepository struct{}

func NewProductServiceRepository() ProductService {
	return &productServiceRepository{}
}

func (r *productServiceRepository) GetProductCart(productIds []uint) ([]domain.ProductCart, error) {

	data := map[string]interface{}{
		"productsList": productIds,
	}

	jsonCartProductReq, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error encoding JSON: %v\n", err)
		return nil, err
	}

	url := config.ENV.URL_PRODUCT_SERVICE + "/cart-product"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonCartProductReq))
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return nil, err
	}

	var productCarts []domain.ProductCart

	err = json.Unmarshal(body, &productCarts)
	if err != nil {
		fmt.Println("Error unmarshalling response data:", err)
		return nil, err

	}

	return productCarts, nil

}

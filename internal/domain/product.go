package domain

type ProductServiceResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Price    int    `json:"price"`
	Stock    int    `json:"stock"`
	Qty      int    `json:"qty"`
	Category struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"category"`
}

type ProductCart struct {
	ID        int    `json:"id"`
	ProductId int    `json:"product_id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
	Qty       int    `json:"qty"`
	Category  struct {
		Name string `json:"name"`
		Slug string `json:"slug"`
	} `json:"category"`
}

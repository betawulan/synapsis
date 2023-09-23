package model

type Product struct {
	ID    int64  `json:"-"`
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

type Category struct {
	ID   int64  `json:"-"`
	Name string `json:"category"`
}

type ProductCategoryFilter struct {
	Category string `json:"category"`
}

type ProductCategory struct {
	ID         int64    `json:"id"`
	ProductID  int64    `json:"product_id"`
	CategoryID int64    `json:"category_id"`
	Product    Product  `json:"product"`
	Category   Category `json:"category"`
}

type ShoppingCart struct {
	ID                int64    `json:"-"`
	UserID            int64    `json:"user_id"`
	ProductCategoryID int64    `json:"product_category_id"`
	Product           Product  `json:"product"`
	Category          Category `json:"category"`
}

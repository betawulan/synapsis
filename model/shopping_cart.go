package model

type ShoppingCart struct {
	ID                int64    `json:"id"`
	UserID            int64    `json:"user_id"`
	ProductCategoryID int64    `json:"product_category_id"`
	Product           Product  `json:"product"`
	Category          Category `json:"category"`
}

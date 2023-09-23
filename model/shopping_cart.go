package model

type ShoppingCart struct {
	ID                int64    `json:"id"  swaggerignore:"true"`
	UserID            int64    `json:"user_id"  swaggerignore:"true"`
	ProductCategoryID int64    `json:"product_category_id"`
	Product           Product  `json:"product"  swaggerignore:"true"`
	Category          Category `json:"category" swaggerignore:"true"`
}

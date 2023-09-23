package service

import (
	"context"

	"github.com/betawulan/synapsis/model"
)

type AuthService interface {
	Register(ctx context.Context, user model.User) error
	Login(ctx context.Context, role string, email string, password string) (string, error)
}

type ProductService interface {
	Fetch(ctx context.Context, filter model.ProductCategoryFilter) ([]model.ProductCategory, error)
}

type ShoppingCartService interface {
	Create(ctx context.Context, tokenString string, shoppingCart model.ShoppingCart) error
	Delete(ctx context.Context, tokenString string, ID int64) error
	Read(ctx context.Context, tokenString string) ([]model.ShoppingCart, error)
}

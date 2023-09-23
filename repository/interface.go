package repository

import (
	"context"

	"github.com/betawulan/synapsis/model"
)

type AuthRepository interface {
	Register(ctx context.Context, user model.User) error
	Login(ctx context.Context, role string, email string, password string) (model.User, error)
}

type ProductRepository interface {
	Fetch(ctx context.Context, filter model.ProductCategoryFilter) ([]model.ProductCategory, error)
}

type ShoppingCartRepository interface {
	Create(ctx context.Context, shoppingCart model.ShoppingCart) (model.ShoppingCart, error)
	Delete(ctx context.Context, ID int64) error
	Read(ctx context.Context, user int64) ([]model.ShoppingCart, error)
}

type TransactionRepository interface {
	CreateTransaction(ctx context.Context, userID int64, productCategoryIDs []int, sumPrice int) error
	SumPrice(ctx context.Context, userID int64, productCategoryIDs []int) (int, error)
}

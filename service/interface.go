package service

import (
	"context"

	"github.com/betawulan/synapsis/model"
)

type AuthService interface {
	Register(ctx context.Context, user model.User) error
	Login(ctx context.Context, role string, email string, password string) (string, error)
}

type OnlineStoreService interface {
	Fetch(ctx context.Context, filter model.ProductCategoryFilter) (model.ProductCategoryResponse, error)
}
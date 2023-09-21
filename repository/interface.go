package repository

import (
	"context"

	"github.com/betawulan/synapsis/model"
)

type AuthRepository interface {
	Register(ctx context.Context, user model.User) error
	Login(ctx context.Context, role string, email string, password string) (model.User, error)
}
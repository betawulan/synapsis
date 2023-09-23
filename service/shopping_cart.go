package service

import (
	"context"

	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/repository"
	"github.com/golang-jwt/jwt"
)

type shoppingCartService struct {
	shoppingCartRepo repository.ShoppingCartRepository
	secretKey        []byte
}

func (s shoppingCartService) Create(ctx context.Context, tokenString string, shoppingCart model.ShoppingCart) error {
	claim := claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	shoppingCart.UserID = claim.ID

	_, err = s.shoppingCartRepo.Create(ctx, shoppingCart)
	if err != nil {
		return err
	}

	return nil
}

func (s shoppingCartService) Delete(ctx context.Context, tokenString string, ID int64) error {
	claim := claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	err = s.shoppingCartRepo.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}

func (s shoppingCartService) Read(ctx context.Context, tokenString string) ([]model.ShoppingCart, error) {
	claim := claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return s.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	shoppingCarts, err := s.shoppingCartRepo.Read(ctx, claim.ID)
	if err != nil {
		return nil, err
	}

	return shoppingCarts, nil
}

func NewShoppingCartService(shoppingCartRepo repository.ShoppingCartRepository, secretKey []byte) ShoppingCartService {
	return shoppingCartService{
		shoppingCartRepo: shoppingCartRepo,
		secretKey:        secretKey,
	}
}

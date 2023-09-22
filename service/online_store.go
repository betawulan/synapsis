package service

import (
	"context"

	"github.com/betawulan/synapsis/model"
	"github.com/betawulan/synapsis/repository"
	"github.com/golang-jwt/jwt"
)

type onlineStoreService struct {
	onlineStoreRepo repository.OnlineStoreRepository
	secretKey       []byte
}

func (o onlineStoreService) Fetch(ctx context.Context, filter model.ProductCategoryFilter) (model.ProductCategoryResponse, error) {
	productCategories, err := o.onlineStoreRepo.Fetch(ctx, filter)
	if err != nil {
		return model.ProductCategoryResponse{}, err
	}

	return model.ProductCategoryResponse{Products: productCategories}, nil
}

func (o onlineStoreService) Create(ctx context.Context, tokenString string, shoppingCart model.ShoppingCart) (model.ShoppingCart, error) {
	claim := claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return o.secretKey, nil
	})
	if err != nil {
		return model.ShoppingCart{}, err
	}

	if !token.Valid {
		return model.ShoppingCart{}, err
	}

	shoppingCart, err = o.onlineStoreRepo.Create(ctx, shoppingCart)
	if err != nil {
		return model.ShoppingCart{}, err
	}

	return shoppingCart, nil
}

func (o onlineStoreService) Delete(ctx context.Context, tokenString string, userID int64, productCategoryID int64) error {
	claim := claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return o.secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	err = o.onlineStoreRepo.Delete(ctx, userID, productCategoryID)
	if err != nil {
		return err
	}

	return nil
}

func NewOnlineStoreService(onlineStoreRepo repository.OnlineStoreRepository, secretKey []byte) OnlineStoreService {
	return onlineStoreService{
		onlineStoreRepo: onlineStoreRepo,
		secretKey:       secretKey,
	}
}

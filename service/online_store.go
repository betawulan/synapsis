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

func (o onlineStoreService) Fetch(ctx context.Context, filter model.ProductCategoryFilter) ([]model.ProductCategory, error) {
	productCategories, err := o.onlineStoreRepo.Fetch(ctx, filter)
	if err != nil {
		return nil, err
	}

	return productCategories, nil
}

func (o onlineStoreService) Create(ctx context.Context, tokenString string, shoppingCart model.ShoppingCart) error {
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

	shoppingCart.UserID = claim.ID

	_, err = o.onlineStoreRepo.Create(ctx, shoppingCart)
	if err != nil {
		return err
	}

	return nil
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

func (o onlineStoreService) Read(ctx context.Context, tokenString string) ([]model.ShoppingCart, error) {
	claim := claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return o.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, err
	}

	shoppingCarts, err := o.onlineStoreRepo.Read(ctx, claim.ID)
	if err != nil {
		return nil, err
	}

	return shoppingCarts, nil
}

func NewOnlineStoreService(onlineStoreRepo repository.OnlineStoreRepository, secretKey []byte) OnlineStoreService {
	return onlineStoreService{
		onlineStoreRepo: onlineStoreRepo,
		secretKey:       secretKey,
	}
}

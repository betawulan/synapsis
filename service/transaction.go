package service

import (
	"context"

	"github.com/betawulan/synapsis/repository"
	"github.com/golang-jwt/jwt"
)

type transactionService struct {
	transactionRepo repository.TransactionRepository
	secretKey       []byte
}

func (t transactionService) Checkout(ctx context.Context, tokenString string, productCategoryIDs []int) error {
	claim := claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(token *jwt.Token) (interface{}, error) {
		return t.secretKey, nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return err
	}

	totalPrice, err := t.transactionRepo.SumPrice(ctx, claim.ID, productCategoryIDs)
	if err != nil {
		return err
	}

	err = t.transactionRepo.CreateTransaction(ctx, claim.ID, productCategoryIDs, totalPrice)
	if err != nil {
		return err
	}

	return nil
}

func NewTransactionService(transactionRepo repository.TransactionRepository, secretKey []byte) TransactionService {
	return transactionService{
		transactionRepo: transactionRepo,
		secretKey:       secretKey,
	}
}
